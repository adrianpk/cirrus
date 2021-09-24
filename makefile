format:
	gofmt -s -w .

build:
	go build ./...

test:
	go clean -testcache
	go test ./... -v
