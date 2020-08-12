PROJECTNAME := $(shell basename "$(PWD)")
GOBASE := $(shell pwd)
GOBIN := $(GOBASE)/bin
MAIN_FILE=main.go

build:
	@echo "  >  Building binary..."
	go build -o ${GOBIN}/${PROJECTNAME} ${MAIN_FILE}

## clean: Clean build files.
clean:
	@-rm $(GOBIN)/$(PROJECTNAME) 2> /dev/null

unit_test:
	go test ./...

run:
	go run ${MAIN_FILE}