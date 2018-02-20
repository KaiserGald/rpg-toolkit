# makefile for unlichtServer
# 7 February 2018
# Code is licensed under the MIT License
# © 2018 Scott Isenberg

# Include configuration
include data/conf/app.cfg
include data/conf/output.cfg
include scripts/lib/make.lib

# Path to mimic
MIMICPATH=$(GOPATH)/src/github.com/KaiserGald/mimic

all: deps test build install clean
	@echo -e $(PURPLE)$(BINARY_NAME)$(NC) successfully installed and started\!

deps:
	@echo -e Grabbing dependencies...
	@go get -u github.com/KaiserGald/logger
	$(DONE)

build:
	@echo -e $(NC)Building $(PURPLE)$(BINARY_NAME)$(NC) server...
	@go build -o $(BINARY_NAME)
	$(DONE)
	@echo -e Writing conf file...
	@echo "{\n\t\"Dir\": \"$(shell pwd)\"\n}" > $(CONFDIR)/conf.json
	$(DONE)

install:
	@echo -e Installing $(PURPLE)$(BINARY_NAME)$(NC) Server...
	@cp -u $(BINARY_NAME) $(INSTALLPATH)
	$(DONE)

# only create directories if they don't already exist
ifeq ("$(wildcard $(WEBINSTALLDIR))", "")
	@mkdir $(WEBINSTALLDIR)
endif

ifeq ("$(wildcard $(APPINSTALLDIR))", "")
	@mkdir $(APPINSTALLDIR)
endif
	$(DONE)

clean:
	@echo -e Cleaning up install.
	@cp -u $(BINARY_NAME) $(DEVBINDIR)
	@rm -f $(BINARY_NAME)
	$(DONE)

test:
	@go test | $(SED_COLORED)

.PHONY: all test install clean
