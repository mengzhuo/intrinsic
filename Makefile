all: scanner sse2 sse3 ssse3 sse41 sse42 avx avx2

scanner:
	rm -rf scanner
	go build scanner.go

%/inst_amd64.go: scanner
	mkdir -p `dirname $@`
	./scanner -out func -feature `dirname $@` > $@ 
	gofmt -w $@

%/inst_amd64.s: scanner
	mkdir -p `dirname $@`
	./scanner -out asm -feature `dirname $@` > $@ 


sse2: sse2/inst_amd64.go sse2/inst_amd64.s
sse3: sse3/inst_amd64.go sse3/inst_amd64.s
ssse3: ssse3/inst_amd64.go ssse3/inst_amd64.s
sse41: sse41/inst_amd64.go sse41/inst_amd64.s
sse42: sse42/inst_amd64.go sse42/inst_amd64.s
avx: avx/inst_amd64.go avx/inst_amd64.s
avx2: avx2/inst_amd64.go avx2/inst_amd64.s

clean: 
	rm */inst_a*


.PHONY: all clean
