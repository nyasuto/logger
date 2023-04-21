
build: test
	go build cmd/sample.go

test:
	go test -v ./...