build:
	go get -d all
	go build -o klingon cmd/main.go

all: build
