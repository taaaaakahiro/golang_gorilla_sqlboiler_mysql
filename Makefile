run:
	go run ./pkg/cmd/main.go

test:
	go test ./...

env:
	direnv allow .

boiler:
	export GOPATH=$HOME/go
    export PATH=$PATH:$GOPATH/bin
	sqlboiler mysql -c sqlboiler/sqlboiler.toml -o models --no-tests