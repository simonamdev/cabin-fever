make:
	modd

test:
	go test ./...

build:
	go build -o cabinserver -ldflags "-linkmode external -extldflags -static" -a main.go