.PHONY: all build run gotool clean help

BINARY="master-gin"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

gotool:
	go fmt ./
	go vet ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - format go file and compile it to binary"
	@echo "make build - compile go file to binary"
	@echo "make run - run go directly"
	@echo "make clean - remove binary file vim swap files"
	@echo "make gotool - run 'fmt' and 'vet'"