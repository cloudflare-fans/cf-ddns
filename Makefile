.PHONY: all
all: clean build

.PHONY: make-target-dir
make-target-dir:
	mkdir -p target/

.PHONY: cross-env
cross-env:
	- docker buildx create --use --name default-cross default

.PHONY: clean-container
clean-container:
	- docker stop cf-ddns && docker rm cf-ddns

.PHONY: build-debian-arm64
build-debian-arm64: make-target-dir cross-env clean-container
	docker buildx build -f Dockerfile_debian_arm64 --platform linux/arm64 -t cf-ddns:latest . --load
	docker run -d --name cf-ddns cf-ddns:latest
	docker cp cf-ddns:/cf-ddns.deb ./target/cf-ddns_1.0.0_arm64.deb
	docker stop cf-ddns && docker rm cf-ddns

.PHONY: build-debian-amd64
build-debian-amd64: make-target-dir cross-env clean-container
	docker buildx build -f Dockerfile_debian_amd64 --platform linux/amd64 -t cf-ddns:latest . --load
	docker run -d --name cf-ddns cf-ddns:latest
	docker cp cf-ddns:/cf-ddns.deb ./target/cf-ddns_1.0.0_amd64.deb
	docker stop cf-ddns
	docker rm cf-ddns

.PHONY: build-macos-arm64
build-macos-arm64:
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ./target/cfddns-darwin-arm64

.PHONY: build-macos-amd64
build-macos-arm64:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./target/cfddns-darwin-amd64

.PHONY: build-windows-arm64
build-windows-arm64:
	CGO_ENABLED=1 GOOS=windows GOARCH=arm64 go build -o ./target/cfddns-windows-arm64.exe

.PHONY: build-windows-amd64
build-windows-amd64:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./target/cfddns-windows-amd64.exe

.PHONY: build-windows-i386
build-windows-i386:
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ./target/cfddns-windows-i386.exe

.PHONY: build
build: build-debian-arm64 build-debian-amd64
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./target/cfddns-darwin-amd64
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o ./target/cfddns-darwin-arm64
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./target/cfddns-windows-amd64.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=386 go build -o ./target/cfddns-windows-i386.exe
	CGO_ENABLED=1 GOOS=windows GOARCH=arm64 go build -o ./target/cfddns-windows-arm64.exe

.PHONY: clean
clean:
	rm -rf ./target/*
