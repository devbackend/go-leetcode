lint-install:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

lint:
	golangci-lint run

test:
	go clean -testcache
	go test ./... --race

bench:
	go test -bench=. ./...

test-coverage:
	go clean -testcache
	go test ./... --race -coverprofile cover.out
	go tool cover -func cover.out