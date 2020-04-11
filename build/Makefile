GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet 
BINARY_NAME=huicd
    
all: test build
build: 
	$(GOVET) ../cmd/huicd
	$(GOBUILD) -o $(BINARY_NAME) -v ../cmd/huicd
test: 
	$(GOTEST) ../cmd/huicd
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ../cmd/huicd
	DATA_DIR=/tmp ./$(BINARY_NAME) server