test: lint
	go test ./... -cover | sort -n -k2.37 | sed 's/github.com\/jackytck\/gowboy\///g'

test-ci:
	go test ./... -cover | sort -n -k2.37 | sed 's/github.com\/jackytck\/gowboy\///g'

testv:
	go test ./... -cover -v

lint:
	golint ./...
	golangci-lint run
