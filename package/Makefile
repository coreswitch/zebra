##
## Makefile for zebra.
##
PACKAGE_NAME:=zebra

ZEBRA_VERSION:=1.8.0

PACKAGE_VERSION:=$(ZEBRA_VERSION)
PACKAGE_LICENSE:=Apache License version 2.0
PACKAGE_DESCRIPTION:=zebra routing software
PACKAGE_MAINTAINER:=ishi@coreswitch.io
PACKAGE_URL:=http://coreswitch.io/
PACKAGE_VENDOR:=coreswitch

OUT_DIR:=out
STAGING_DIR:=$(OUT_DIR)/stage

.PHONY: package stage clean files

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
	--deb-changelog ../../ChangeLog \
	--deb-no-default-config-files \
	-t deb .
	touch $(OUT_DIR)/package.timestamp

stage: files zebra

files: $(STAGING_DIR)
	rsync -avz ./files/ubuntu/ $(STAGING_DIR)/

zebra: $(STAGING_DIR)/usr/bin
	go get github.com/coreswitch/zebra/rib/ribd
	cp $(GOPATH)/bin/ribd $(STAGING_DIR)/usr/bin

$(STAGING_DIR):
	mkdir -p $(STAGING_DIR)

$(STAGING_DIR)/usr/bin:
	mkdir -p $(STAGING_DIR)/usr/bin

clean:
	rm -rf $(OUT_DIR)
