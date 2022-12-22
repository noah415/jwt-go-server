run:
	go run ./cmd/main.go

start:
	./bin/main

build:
	go build -o bin/main cmd/main.go

clean:
	rm -rf bin