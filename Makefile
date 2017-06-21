scanner:
	rm scanner
	go build scanner.go

%/inst_amd64.go: scanner
	mkdir -p `dirname $@`
	./scanner -out func > $@ && gofmt -w $@

%/inst_amd64.s: scanner
	mkdir -p `dirname $@`
	./scanner -out asm > $@ 


sse2: sse2/*
sse3: sse3/inst_*
ssse3: ssse3/*
#sse41: scanner sse41/inst_*
#sse42: scanner sse42/*
#avx: scanner avx/*
#avx2: scanner avx2/*


clean: 
	rm */inst_a*

all: scanner sse2 sse3 ssse3 #sse41 sse42 avx avx2
	echo "all"

.PHONY: all clean
