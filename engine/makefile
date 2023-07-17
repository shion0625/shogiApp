test:
	go test ./...

test-log:
	go test -v ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out


test-coverage-html:
	go tool cover -html=coverage.out
