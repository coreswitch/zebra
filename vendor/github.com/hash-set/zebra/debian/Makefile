##
## Makefile for Zebra package.
##
#GOPATH:=$(shell pwd)/go
PACKAGE_NAME:=ribd

RIBD_VERSION:=2.0
ESI_VERSION:=77

PACKAGE_VERSION:=$(RIBD_VERSION)-$(ESI_VERSION)
PACKAGE_LICENSE:=Apache License version 2.0
PACKAGE_DESCRIPTION:=ribd for ESI network controller
PACKAGE_MAINTAINER:=support@ntti3.com
PACKAGE_URL:=http://www.ntti3.com/about
PACKAGE_VENDOR:=ntti3.com

OUT_DIR:=out
STAGING_DIR:=$(OUT_DIR)/stage

.PHONY: package stage clean

package: $(OUT_DIR)/package.timestamp

$(OUT_DIR)/package.timestamp: stage
	@echo "Creating package"
	@rm -f $(OUT_DIR)/$(PACKAGE_NAME)_$(PACKAGE_VERSION)*.deb
	@cd $(OUT_DIR) && \
	fpm -s dir -C $(realpath $(STAGING_DIR)) \
	--name $(PACKAGE_NAME) \
	--version $(PACKAGE_VERSION) \
	--license "$(PACKAGE_LICENSE)" \
	--maintainer "$(PACKAGE_MAINTAINER)" \
	--description "$(PACKAGE_DESCRIPTION)" \
  --url "$(PACKAGE_URL)" \
	--vendor "$(PACKAGE_VENDOR)" \
	--before-install ../scripts/before_install.sh \
	--after-install ../scripts/after_install.sh \
	--before-remove ../scripts/before_remove.sh \
	--after-remove ../scripts/after_remove.sh \
	--deb-changelog ../ChangeLog \
	--deb-no-default-config-files \
	-t deb .
	touch $(OUT_DIR)/package.timestamp

stage: static ribd

static: $(STAGING_DIR)
	rsync -avz ./packaging/ubuntu/ $(STAGING_DIR)/

ribd: $(STAGING_DIR)/usr/bin
	go get github.com/hash-set/zebra/rib/ribd
	cp $(GOPATH)/bin/ribd $(STAGING_DIR)/usr/bin
	(cd ../vrf; make)
	cp ../vrf/vrf_socket.so $(STAGING_DIR)/usr/bin
	cp ../tools/gobgp-vrf.sh $(STAGING_DIR)/usr/bin

$(STAGING_DIR):
	mkdir -p $(STAGING_DIR)

$(STAGING_DIR)/usr/bin:
	mkdir -p $(STAGING_DIR)/usr/bin

clean:
	rm -rf $(OUT_DIR)/*
