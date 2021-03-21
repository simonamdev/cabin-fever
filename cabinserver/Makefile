make:
	modd

test:
	go test ./...

build:
	rm -rf cabinserver
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cabinserver .
