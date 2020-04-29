.PHONY: build
build:
	go build -v
	./hh -config=config.json
.DEFAULT_GOAL := build