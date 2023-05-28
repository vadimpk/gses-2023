build:
	go build -o ./.bin/srv cmd/main.go

run: build
	./.bin/srv
