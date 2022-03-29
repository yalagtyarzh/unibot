.PHONY:
.SILENT:

build:
	go build -o ./.bin/unibot cmd/app/main.go

run: build
	./.bin/unibot
