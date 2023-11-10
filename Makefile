build:
	GOOS=linux GOARCH=amd64 go build -o bin/main main.go

run:
	go run main.go

test:
	cd test && go test -race ./... && cd..

config:
	cp .env.example .env

tidy:
	go mod tidy
key:
	go run zoro.go key:generate