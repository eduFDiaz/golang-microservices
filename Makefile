executableName=testApi
clean:
	echo "stoping if running and cleaning"
	rm -rf ./bin
	killall $(executableName) || true
test:
	echo "Testing..."
	go test -coverprofile cp.out ./mvc/...
	go tool cover -html=cp.out
build:
	echo "Building..."
	go build -o bin/$(executableName) mvc/main.go
run:
	echo "Running..."
	./bin/$(executableName) &
all: test build run