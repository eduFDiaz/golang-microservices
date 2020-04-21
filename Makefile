test:
    echo "Testing..."
	go test -coverprofile cp.out ./mvc/...
	go tool cover -html=cp.out
build:
	echo "Building..."
	go build -o bin/main mvc/main.go
run:
	echo "Running..."
	./bin/main
all: test build run