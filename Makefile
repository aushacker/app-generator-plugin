#PWD := $(shell pwd)
#LOCAL_BIN ?= $(PWD)/bin

############################################################
# clean section
############################################################

.PHONY: clean
clean:
	-rm appGenerator
	-podman rmi localhost/appgenerator

############################################################
# build section
############################################################

.PHONY: build
build:
	go build ./cmd/appGenerator
	podman build -t localhost/appgenerator .

############################################################
# tests section
############################################################

.PHONY: test
test:
	kustomize build --enable-alpha-plugins test/01-minimal/