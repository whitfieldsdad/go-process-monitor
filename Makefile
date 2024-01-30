default: clean build

NAME = go-process-monitor
LDFLAGS = "-s -w"

help:
	@echo "make [test]"

build: windows darwin linux

clean:
	rm -rf bin

windows:
	GOOS=windows GOARCH=amd64 go build -o bin/${NAME}-windows-amd64.exe main.go
	GOOS=windows GOARCH=arm64 go build -o bin/${NAME}-windows-arm64.exe main.go

darwin:
	GOOS=darwin GOARCH=amd64 go build -o bin/${NAME}-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/${NAME}-darwin-arm64 main.go

linux:
	GOOS=linux GOARCH=amd64 go build -o bin/${NAME}-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/${NAME}-linux-arm64 main.go

test:
	go test -v .
