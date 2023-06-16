
compile:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go

build:
	go build -o bin/main main.go

run:
	go run main.go

test_banana:
	go test banana/banana_test.go -v

test_hello:
	go test hello/hello_test.go -v

test_date:
	go test date/date_test.go -v

test_all:
	go test ./... -v