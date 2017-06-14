.PHONY: all
	
all : sse2 sse3 ssse3 sse41 sse42 

sse2:
	mkdir -p sse2
	go run scanner.go  -out func > sse2/inst_amd64.go && gofmt -w sse2/inst_amd64.go
	go run scanner.go  -out asm > sse2/inst_amd64.s 

sse3:
	mkdir -p sse3
	go run scanner.go -feature sse3 -out func > sse3/inst_amd64.go && gofmt -w sse3/inst_amd64.go
	go run scanner.go -feature sse3 -out asm > sse3/inst_amd64.s

ssse3:
	mkdir -p ssse3
	go run scanner.go -feature ssse3 -out func > ssse3/inst_amd64.go
	go run scanner.go -feature ssse3 -out asm > ssse3/inst_amd64.s

sse41:
	mkdir -p sse41
	go run scanner.go -feature sse41 -out func > sse41/inst_amd64.go
	go run scanner.go -feature sse41 -out asm > sse41/inst_amd64.s

sse42:
	mkdir -p sse42
	go run scanner.go -feature sse42 -out func > sse42/inst_amd64.go
	go run scanner.go -feature sse42 -out asm > sse42/inst_amd64.s
