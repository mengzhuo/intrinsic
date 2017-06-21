feature_list := sse2 sse3 ssse3 sse41 sse42 avx avx2

scanner:
	rm -rf .tmp_scanner
	go build  -o .tmp_scanner  scanner.go

define gen
	mkdir -p $1
	./.tmp_scanner -out func -feature $1 > $1/inst_amd64.go
	gofmt -w $1

	mkdir -p $1
	./.tmp_scanner -out asm -feature $1 > $1/inst_amd64.s

endef

define testfunc
	cd $1
	go test -v
endef

all: scanner
	$(foreach feature, $(feature_list), $(call gen, $(feature)))

test: 
	$(foreach feature, $(feature_list), $(call testfunc, $(feature)))

clean: 
	rm .tmp_scanner
	rm */inst_a*	

.PHONY: all clean test
