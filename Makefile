
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/main-amd64-linux main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386.exe main.go
	GOOS=windows GOARCH=amd64 go build -o bin/main-indows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/main-amd64-darwin main.go
	GOOS=js GOARCH=wasm go build -o bin/main-wasm.wasm main.go

build:
	go build -o bin/main main.go

run:
	go run main.go

git_push:
	git add . && git commit -m "update github.com/stretchr/testify/assert" && git push

test_banana:
	go test banana/banana_test.go -v

test_hello:
	go test hello/hello_test.go -v

test_date:
	go test date/date_test.go -v

test_all:
	go test ./... -v