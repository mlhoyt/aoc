all: vet test

vet:
	go vet ./...

test:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -func coverage.out
