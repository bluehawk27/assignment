SHELL := /bin/bash
BINARY=assignmentd
TESTDIRS=`go list ./... | grep -v /vendor/`
BASEDIR=$(shell pwd)

# create binary for Docker
default:
	rm -rf build/
	mkdir build
	CGO_ENABLED=0 GOOS=linux go build -o build/$(BINARY)

deps:
	glide install

test:
	PROXY_CONFIGPATH=$(BASEDIR) PROXY_CONFIGFILE=config-test go test --cover -v $(TESTDIRS)

start: 
	go install
	assignment

clean:
	rm -rf build/


