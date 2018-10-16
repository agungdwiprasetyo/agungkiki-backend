.PHONY : build test

build: main.go 
	go build

test: ./token 
	go test -count=1 -race ./token