.PHONY: build
build:
	GOOS=linux GOARCH=amd64 go build -o rename-cli-linux
	GOOS=linux GOARCH=arm64 go build -o rename-cli-linux-arm64
	GOOS=windows GOARCH=amd64 go build -o rename-cli.exe
	GOOS=windows GOARCH=arm64 go build -o rename-cli-arm64.exe
	GOOS=android GOARCH=arm64 go build -o rename-cli-android
