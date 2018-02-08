# makefile for unlichtServer
# 7 February 2018
# Code is licensed under the MIT License
# © 2018 Scott Isenberg

# Variable declarations

# To change the name of the app simply change the BINARY_NAME variable

# Name of the app to be installed
BINARY_NAME=unlichtServer

# Path to the server binary install
INSTALLPATH=/usr/local/bin/
BINPATH=$(INSTALLPATH)$(BINARY_NAME)
ARGS=-d -c

# Log path
LOGPATH=/var/log

# Dev bin path
DEVBINDIR=bin
DEVBINPATH=$(DEVBINDIR)/$(BINARY_NAME)

# Installation paths
WEBINSTALLDIR=/srv/$(BINARY_NAME)
APPINSTALLDIR=$(WEBINSTALLDIR)/app
ASSETINSTALLDIR=$(APPINSTALLDIR)/assets
CONFINSTALLDIR=$(APPINSTALLDIR)/conf

# Source paths
APPDIR=./app
COMPDIR=$(APPDIR)/components
STATICDIR=$(APPDIR)/static
ASSETSDIR=$(APPDIR)/assets
CSSDIR=$(ASSETSDIR)/css
IMGDIR=$(ASSETSDIR)/img
JSDIR=$(ASSETSDIR)/js
CONFDIR=./data/conf

# Service paths
LOCALSERVICEDIR=/lib/systemd/system
LOCALSERVICECONFIG=/$(BINARY_NAME).service
SERVICEDIR=data/systemd
SERVICECONFIG=$(SERVICEDIR)/$(BINARY_NAME).service

# Path to mimic
MIMICPATH=$(GOPATH)/src/github.com/KaiserGald/mimic

# Color variables
RED='\033[0;31m'
GREEN='\033[0;32m'
WHITE='\033[1;37m'
PURPLE='\e[95m'
CYAN='\e[36m'
YELLOW='\033[1;33m'
ORANGE='\033[38;5;208m'
NC='\033[0m'

# script to colorize the go test output
SED_COLORED=sed ''/'\(--- PASS\)'/s//$$(printf $(GREEN)---\\x20PASS)/'' | sed ''/PASS/s//$$(printf $(GREEN)PASS)/'' | sed  ''/'\(=== RUN\)'/s//$$(printf $(YELLOW)===\\x20RUN)/'' | sed ''/ok/s//$$(printf $(GREEN)ok)/'' | sed  ''/'\(--- FAIL\)'/s//$$(printf $(RED)---\\x20FAIL)/'' | sed  ''/FAIL/s//$$(printf $(RED)FAIL)/'' | sed ''/RUN/s//$$(printf $(YELLOW)RUN)/'' | sed ''/?/s//$$(printf $(ORANGE)?)/'' | sed ''/'\(^\)'/s//$$(printf $(NC))/''

# script to see if service is running
ISSERVICERUNNING=$(shell ps -ef | grep -v grep | grep $(BINARY_NAME) | wc -l)

# script to make the .service file
MAKESERVICECONFIG=@cat $(SERVICEDIR)/srvtmpl.service | sed ''s/'DESCRIPTION'/'$(BINARY_NAME) Server'/'' | sed ''s~'CONDITIONPATHEXISTS'~$(BINPATH)~'' | sed ''s~'WORKINGDIRECTORY'~$(INSTALLPATH)~'' | sed ''s~'EXECSTART'~$(BINPATH)~'' | sed ''s~'EXSTARTPRE'~$(LOGPATH)/$(BINARY_NAME)~'' | sed ''s~'SYSLOGIDENTIFIER'~$(BINARY_NAME)~'' | sed ''s~'ARGS'~'$(ARGS)'~ > $(SERVICECONFIG)

# done script
DONE=@echo -e $(GREEN)Done\!$(NC)

all: stop deps test build install clean run
	@echo -e $(PURPLE)$(BINARY_NAME)$(NC) successfully installed and started\!
	@echo -e $(ARGS)

stop:
	@echo -e Checking to see if $(PURPLE)$(BINARY_NAME)$(NC) server is running...
ifneq (${ISSERVICERUNNING}, 0)
	@echo -e $(PURPLE)$(BINARY_NAME)$(NC) server is running. Stopping it now.
	@sudo systemctl stop $(BINARY_NAME)
	$(DONE)
else
	@echo -e $(PURPLE)$(BINARY_NAME)$(NC) isn\'t currently running.
endif

deps:
	@echo -e Grabbing dependencies...
	@go get github.com/KaiserGald/mimic
	@go get github.com/KaiserGald/logger
	$(DONE)

build:
	@echo -e $(NC)Building $(PURPLE)$(BINARY_NAME)$(NC) server...
	@go build -o $(BINARY_NAME)
	$(DONE)
	@echo -e Writing conf file...
	@echo -e "{\n\t\"Dir\": \"$(shell pwd)\"\n}" > $(CONFDIR)/conf.json
	$(DONE)

install:
	@echo -e Installing $(PURPLE)$(BINARY_NAME)$(NC) Server...
	@sudo cp -u $(BINARY_NAME) $(INSTALLPATH)
	$(DONE)
	@echo -e Installing $(PURPLE)Mimic$(NC)...
	@make -C $(MIMICPATH)
	@printf '#!/bin/bash\nsudo journalctl -f -u $(BINARY_NAME) -o cat' > $(BINARY_NAME)_log
	@sudo chmod +x $(BINARY_NAME)_log
	$(DONE)

ifeq ("$(wildcard $(WEBINSTALLDIR))", "")
	@sudo mkdir $(WEBINSTALLDIR)
endif

ifeq ("$(wildcard $(APPINSTALLDIR))", "")
	@sudo mkdir $(APPINSTALLDIR)
endif

ifeq ("$(wildcard $(LOCALSERVICECONFIG))", "")
	@sudo rm -f $(LOCALSERVICECONFIG)
endif

ifneq ("$(wildcard $(SERVICECONFIG))", "")
	@rm -f $(SERVICECONFIG)
endif
	${MAKESERVICECONFIG}

	@sudo cp -u $(SERVICECONFIG) $(LOCALSERVICEDIR)
	$(DONE)

clean:
	@echo -e Cleaning up install.
	@cp -u $(BINARY_NAME) $(DEVBINDIR)
	@rm -f $(BINARY_NAME)
	$(DONE)

run:
	@echo -e Starting up $(PURPLE)$(BINARY_NAME)$(NC) server as a daemon.
	@sudo systemctl start $(BINARY_NAME)
	$(DONE)
	@echo -e Starting up $(PURPLE)mimic$(NC).
	@sudo mimic -c -w "app:$(APPINSTALLDIR)" &
	$(DONE)
	@echo -e Type $(WHITE)\'$(CYAN)./$(BINARY_NAME)_log$(WHITE)\'$(NC) to open the server log in the terminal.

test:
	@go test | $(SED_COLORED)
