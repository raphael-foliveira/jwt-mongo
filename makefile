build-osx:
	GOOS=darwin GOARCH=amd64 go build -o bin/app.darwin -v ./cmd/app

build-linux:
	GOOS=linux GOARCH=amd64 go build -o bin/app -v ./cmd/app

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/app.exe -v ./cmd/app

build-all: build-osx build-linux build-windows

run: build-linux
	./bin/app

dev:
	air