install-lint:
	go get -u golang.org/x/lint/golint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.21.0

lint:
	golint ./...
	golangci-lint run

test:
	go test ./... -cover | sort -n -k2.37 | sed 's/github.com\/jackytck\/gowboy\///g'

testv:
	go test ./... -cover -v
