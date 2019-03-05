ifndef BUILD_NUMBER
BUILD_NUMBER = $(shell date +%Y%m%d%H%M)
endif

VERSION=1.0.$(BUILD_NUMBER)
CONFIG_DIR = $(CURDIR)/config
PRIVATE_FILE = $(CURDIR)/local/private_data.yml
GO_SERVER_PID=$(shell pidof roybot)

.PHONY: echoenv clean shutdown run

all: echoenv recovery setconfig shutdown run

echoenv:
	@echo "*** Environment ***"
	@echo "Version : $(VERSION)"
	@echo "CONFIG_DIR : $(CONFIG_DIR)"

recovery:
	@echo "*** Reconvery ***"
	@cp $(CURDIR)/local/config_original.yml $(CONFIG_DIR)/config.yml

setconfig:
	@echo "*** Setconfig ***"
	@sed -i "s/@RELEASE_VERSION@/$(VERSION)/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@RELEASE_TIME@/`date +\"%Y-%m-%d %T\"`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@DB_USER@/`cat $(PRIVATE_FILE) | sed '1!d'`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@DB_PASSWORD@/`cat $(PRIVATE_FILE) | sed '2!d'`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@DB_HOST@/`cat $(PRIVATE_FILE) | sed '3!d'`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@LINE_ADMINID@/`cat $(PRIVATE_FILE) | sed '4!d'`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@LINE_SECRET@/`cat $(PRIVATE_FILE) | sed '5!d'`/g" $(CONFIG_DIR)/config.yml
	@sed -i "s/@LINE_TOKEN@/`cat $(PRIVATE_FILE) | sed '6!d'`/g" $(CONFIG_DIR)/config.yml

shutdown:
	@echo "*** Shutdown Server ***"
	@kill -9 $(GO_SERVER_PID) | true

run:
	@echo "*** Run Server ***"
	@go build
	@screen -d -m ./roybot
