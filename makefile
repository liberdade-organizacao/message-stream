default: test

build:
	go build -o main.exe main/main.go

test: build
	./main.exe &
	go test test/api_test.go
	fuser -k 8000/tcp

database:
	# docker pull postgres
	docker run -p 5432:5432 --name psql -e POSTGRES_PASSWORD=mysecretpassword -d postgres
