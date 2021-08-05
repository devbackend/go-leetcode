lint-install:
	@curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin

lint:
	golangci-lint run

test:
	go clean -testcache
	go test ./... --race

test-coverage-generate:
	go test ./... -covermode=atomic -coverprofile=coverage.out -coverpkg ./internal/...
	go test ./... -covermode=atomic -coverprofile=coverage2.tmp -coverpkg ./pkg/...
	tail -n +2 coverage2.tmp >> coverage.out && rm coverage2.tmp

test-coverage:
	make test-coverage-generate
	go tool cover -func=coverage.out

test-coverage-html:
	make test-coverage-generate
	go tool cover -html=coverage.out