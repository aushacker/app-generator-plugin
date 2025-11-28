#PWD := $(shell pwd)
#LOCAL_BIN ?= $(PWD)/bin

############################################################
# clean section
############################################################

.PHONY: clean
clean:
	-rm appGenerator

############################################################
# build section
############################################################

.PHONY: build
build:
	go build ./cmd/appGenerator
