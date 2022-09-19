run:
	go run ./pkg/cmd/main.go

test:
	go test ./...

env:
	direnv allow .
