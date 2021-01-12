default: test

build:
	go build -o main.exe main/main.go

test: build
	./main.exe &
	go test test/api_test.go
	fuser -k 8000/tcp
