all: scanner
	$(foreach feature, $(feature_list), $(call gen, $(feature)))

scanner:
	@rm -rf .tmp_scanner
	@go build  -o .tmp_scanner  scanner.go

define gen
	@mkdir -p $1
	@./.tmp_scanner -out func -feature $1 > $1/inst_amd64.go
	@gofmt -w $1/inst_amd64.go

	@./.tmp_scanner -out asm -feature $1 > $1/inst_amd64.s
	@./.tmp_scanner -out test -feature $1 > $1/inst_test.go
	@gofmt -w $1/inst_test.go

endef

define testfunc
	cd $1 && go test -v ; cd -;
endef


feature_list := sse2 sse3 ssse3 sse41 sse42 #avx avx2

test: 
	$(foreach feature, $(feature_list), $(call testfunc, $(feature)))

clean: 
	rm .tmp_scanner
	rm */inst_a*	

.PHONY: all clean test
