
compile:
	echo "Compiling ..."
	GOOS=linux GOARCH=386 go build -o bin/go-tutorial-linux-386 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/go-tutorial-amd64-linux main.go
	GOOS=windows GOARCH=386 go build -o bin/go-tutorial-windows-386.exe main.go
	GOOS=windows GOARCH=amd64 go build -o bin/go-tutorial-indows-amd64.exe main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/go-tutorial-amd64-darwin main.go
	GOOS=js GOARCH=wasm go build -o bin/go-tutorial-wasm.wasm main.go

build:
	echo "Building ...."
	go build -o bin/go-tutorial main.go

clean:
	echo "Cleaning ...."
	go clean
	rm bin/go-tutorial
	rm bin/go-tutorial-linux-386
	rm bin/go-tutorial-amd64-linux
	rm bin/go-tutorial-windows-386.exe
	rm bin/go-tutorial-amd64-darwin
	rm bin/go-tutorial-wasm.wasm

docker_build:
	docker build -t go-tutorial -f Dockerfile

docker_deploy:
	docker run -d --name go-tutorial go-tutorial

run:
	go run main.go

preview:
	./bin/go-tutorial

git_push:
	git add . && git commit -m "update test_hello struct" && git push

test_banana:
	go test banana/banana_test.go -v

test_hello:
	go test hello/hello_test.go -v

test_date:
	go test date/date_test.go -v

test_all:
	go test ./... -v