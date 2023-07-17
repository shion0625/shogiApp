test:
	go test ./tests/...

test-log:
	go test -v ./tests/...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out


test-coverage-html:
	go tool cover -html=coverage.out
