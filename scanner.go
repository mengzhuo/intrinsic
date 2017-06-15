//+build ignore

package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
)

var (
	csvPath      = flag.String("csv", "x86.csv", "csv file path")
	destPath     = flag.String("desc", "x86desc.csv", "csv file path")
	outPath      = flag.String("out", "func", "output file type")
	printFeature = flag.String("feature", "sse2", "print feature")
	plaform      = flag.String("platform", "amd64", "")
	table        = []*Inst{}
	featureMap   = map[string][]*Inst{}
	descMap      = map[string]string{}
)

type Inst struct {
	Mnemonic, Encoding, Valid32, Valid64, Feature, Tags string

	Args     []string
	FuncName string
	Register string
}

func ParseInst(l []string) (i *Inst) {

	if strings.Contains(l[5], "pseudo") {
		return
	}

	i = &Inst{
		Mnemonic: l[0],
		Encoding: l[1],
		Valid32:  l[2],
		Valid64:  l[3],
		Feature:  featureParse(l[4]),
		Tags:     l[5]}

	ol := strings.Split(i.Mnemonic, " ")

	i.FuncName = ol[0]
	i.Args = ol[1:]
	for j, arg := range i.Args {
		arg = strings.Replace(arg, ",", "", 1)
		if slashI := strings.Index(arg, "/"); slashI != -1 {
			arg, i.Register = arg[:slashI], arg[slashI+1:]
		}
		t := arg
		switch t {
		case "xmm1":
			t = "X1"
		case "xmm":
			t = "X1"
		case "xmm2":
			t = "X2"
		case "mm1":
			t = "M1"
		case "mm2":
			t = "M2"
		case "r32":
			t = "CX"
		case "ymm1":
			t = "Y1"
		case "ymm2":
			t = "Y2"
		}
		i.Args[j] = t
	}
	sort.Reverse(sort.StringSlice(i.Args))
	return
}

func (i *Inst) String() string {
	if i.IsPacked() {
		return fmt.Sprintf("%10s %4v %4v %v", i.FuncName, i.Args, i.Target())
	}
	return fmt.Sprintf("%s %v %s", i.FuncName, i.Args, i.Encoding)
}

func (i *Inst) CovertArgs(t string) string {

	args := make([]string, len(i.Args))
	copy(args, i.Args)

	for i, k := range args {
		switch k {
		case "X1", "X2", "M1", "M2", "Y1", "Y2":
			k = k + " []" + t
		case "r32":
			k = "r32 int32"
		case "imm8u":
			k = "imm8u uint8"
		}
		args[i] = k
	}
	return strings.Join(args, ",  ")
}

func (i *Inst) TrueOpcode() string {
	data := strings.Split(i.Encoding, " ")
	for i, v := range data {
		v = strings.TrimSpace(v)
		switch v {
		case "/r":
			v = ""
		default:
			v = "BYTE $0x" + v + ";"
		}
		data[i] = v
	}
	return strings.Join(data, " ")
}

// the frame must be aligned as Go required
func (i *Inst) FrameSize() (s int) {

	for _, a := range i.Args {
		switch a {
		case "X1", "X2", "M1", "M2", "Y1", "Y2":
			s += 24 // slice size 24
		case "r32":
			s += 4
		case "imm8u":
			s += 1
		}
	}

	if s%8 == 0 {
		return s
	}
	return (s/8 + 1) * 8
}

func (i *Inst) IsPacked() bool {
	return i.Mnemonic[0] == 'P'
}

func (i *Inst) Description() string {
	if s, ok := descMap[i.FuncName]; ok {
		return s
	}
	return ""
}

func (i *Inst) Target() (t []string) {

	t = []string{`byte`}
	last := i.FuncName[len(i.FuncName)-1]
	dst := i.Description()
	if strings.Index(dst, "Integers") != -1 {
		if strings.Index(dst, "Unsigned") != -1 {
			if last == 'B' {
				t = append(t, "uint8")
			}
			if last == 'W' {
				t = append(t, "uint16")
			}
			if last == 'D' {
				t = append(t, "uint32")
			}
			if last == 'Q' {
				t = append(t, "int64")
			}
		} else {
			if last == 'B' {
				t = append(t, "int8")
			}
			if last == 'W' {
				t = append(t, "int16")
			}
			if last == 'D' {
				t = append(t, "int32")
			}
			if last == 'Q' {
				t = append(t, "int64")
			}
		}
		return
	}

	if strings.Index(dst, "Single-Precision Floating-Point") != -1 {
		t = append(t, "float32")
		return
	}
	if strings.Index(dst, "Single-FP") != -1 {
		t = append(t, "float32")
		return
	}

	if strings.Index(dst, "Double-Precision Floating-Point") != -1 {
		t = append(t, "float64")
		return
	}

	if strings.Index(dst, "Double-FP") != -1 {
		t = append(t, "float64")
		return
	}

	return
}

func main() {

	flag.Parse()

	readDesc()

	f, err := os.Open(*csvPath)
	if err != nil {
		log.Fatal(err)
	}
	rdr := csv.NewReader(f)

	records, err := rdr.ReadAll()
	if err != nil {
		log.Print(err)
	}
	for i, l := range records {
		if len(l) < 6 {
			log.Printf("line[%d], less than 6 items got=%s", i, l)
			continue
		}

		inst := ParseInst(l)
		if inst == nil {
			continue
		}
		table = append(table, inst)
		featureMap[inst.Feature] = append(featureMap[inst.Feature], inst)
	}

	if fm, ok := featureMap[*printFeature]; ok {
		makeInst(*printFeature, fm)
	} else {
		log.Fatalf("Feature:%s not existed", *printFeature)
	}
}

func readDesc() {
	f, err := os.Open(*destPath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {

		index := strings.Index(scanner.Text(), "\t")
		if index <= 0 {
			log.Printf("text=%s, len=%d", scanner.Text(), index)
			continue
		}
		k := strings.TrimSpace(scanner.Text()[:index])
		dsr := strings.TrimSpace(scanner.Text()[index:])
		for _, item := range strings.Split(k, "/") {
			descMap[item] = dsr
		}
	}
}

func featureParse(s string) string {
	return strings.ToLower(strings.Replace(s, "_", "", 1))
}

type Data struct {
	Platform    string
	FeatureName string
	InstList    []*Inst
}

var skipInst = map[string]int{
	"MOV": 1,
	"CVT": 1,
}

func makeInst(feature string, instList []*Inst) {

	argsTable := map[string]int{}
	skipedInst := []*Inst{}
	var gen int
	for _, inst := range instList {
		if _, ok := skipInst[inst.FuncName[:3]]; ok {
			continue
		}

		if inst.Valid64 == "V" {
			argsTable[fmt.Sprintf("%v", inst.Args)] += 1
		}
		switch fmt.Sprint(inst.Args) {
		case "[X1 X2]":
		case "[Y1 Y2]":
		default:
			continue
		}
		gen += 1
		skipedInst = append(skipedInst, inst)
	}
	log.Printf("%s gen=%d, skip=%d, ratio=%0.2f%%", feature, gen, len(instList)-gen,
		float64(gen)/float64(len(instList))*100)
	dat := &Data{*plaform, feature, skipedInst}
	tmpl := funcTmpl
	switch *outPath {
	case "asm":
		tmpl = asmTmpl
	case "test":
		tmpl = testTmpl
	}
	tt, err := template.New("root").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = tt.Execute(os.Stdout, dat)
	if err != nil {
		log.Fatal(err)
	}
}

var stepMap = map[string]int{
	"X1":    24,
	"X2":    24,
	"M2":    24,
	"M1":    24,
	"imm8u": 1,
	"r32":   4,
}

func (i *Inst) ArgsToAsm() string {
	buf := []string{}
	switch fmt.Sprintf("%v", i.Args) {
	case "[X1 X2]":
		if _, in := unrecognized[i.FuncName]; in {
			return fmt.Sprintf(X1X2Raw, i.TrueOpcode())
		}
		return fmt.Sprintf(X1X2, i.FuncName)

	case "[Y1 Y2]":
		if _, in := unrecognized[i.FuncName]; in {
			return fmt.Sprintf(Y1Y2Raw, i.TrueOpcode())
		}
		return fmt.Sprintf(Y1Y2, i.FuncName)
	default:
		log.Printf("unsupported args, %s", i.FuncName)
	}
	return strings.Join(buf, "\n\t")
}

var unrecognized = set(`
PACKSSDW
PCMPEQD
PCMPGTD
PMADDWD
PMADDWD
PMULUDQ
PMULUDQ
PSLLD
PSRAD
PSRLD
PSUBD
PUNPCKHDQ
PUNPCKHWD
PUNPCKLDQ
PUNPCKLWD
`)

func set(s string) (r map[string]struct{}) {
	r = map[string]struct{}{}
	for _, n := range strings.Split(s, "\n") {
		n = strings.TrimSpace(n)
		if n == "" {
			continue
		}
		if n[:2] == "//" {
			continue
		}
		r[n] = struct{}{}
	}
	return
}

const funcTmpl = `package {{.FeatureName}}

{{ range $index, $inst := .InstList }}
{{ range $target := .Target }}
// go:noescape

// {{$inst.Description}}
func {{$inst.FuncName}}{{$inst.Register}}{{$target}}({{$inst.CovertArgs $target}})
{{end}}{{end}}
`

const testTmpl = ``

const X1X2 = `FPTOX1X2
	%s X2, X1
	RETX1X2
	`

const X1X2Raw = `FPTOX1X2
	%s BYTE $0xca // $0xca = X2, X1
	RETX1X2
	`
const Y1Y2 = `
	MOVQ a+0(FP), SI
	MOVQ b+24(FP), DI
	MOVOU (SI), Y1
	MOVOU (DI), Y2
	%s Y2, Y1
	MOVOU Y1, (SI)
	RET
	`
const Y1Y2Raw = `
	MOVQ a+0(FP), SI
	MOVQ b+24(FP), DI
	MOVOU (SI), Y1
	MOVOU (DI), Y2
	%s BYTE $0xca // $0xca = X2, X1
	MOVOU Y1, (SI)
	RET
	`
const asmTmpl = `#include "textflag.h"

#define FPTOX1X2 \
	MOVQ a+0(FP), SI;\
	MOVQ b+24(FP), DI;\
	MOVOU (SI), X1;\
	MOVOU (DI), X2;\

#define RETX1X2 \
	MOVOU X1, (SI);\
	MOVOU X2, (DI);\
	RET;\

{{ range $index, $inst := .InstList }}
{{ range $target := .Target }}
TEXT ·{{$inst.FuncName}}{{$inst.Register}}{{$target}}(SB),NOSPLIT,$0-{{$inst.FrameSize}}
	{{$inst.ArgsToAsm}}
{{end}}{{end}}

`
