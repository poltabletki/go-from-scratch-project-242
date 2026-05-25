build:
	go build -o hexlet-path-size ./cmd/hexlet-path-size

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix		

test:
	go test -v ./tests