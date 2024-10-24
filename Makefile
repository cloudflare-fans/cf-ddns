.PHONY: all
	all: clean build

.PHONY: build
build:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./target/cfddns-darwin-amd64
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ./target/cfddns-darwin-arm64
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./target/cfddns-windows-amd64.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ./target/cfddns-windows-i386.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=arm64 go build -o ./target/cfddns-windows-arm64.exe

.PHONY: clean
clean:
	rm -rf ./target/*
