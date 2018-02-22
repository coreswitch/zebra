##
## Makefile for openconfigd package.
##
PACKAGE_NAME:=openconfigd

OPENCONFIGD_VERSION:=0.8.1

PACKAGE_VERSION:=$(OPENCONFIGD_VERSION)
PACKAGE_LICENSE:=Apache License version 2.0
PACKAGE_DESCRIPTION:=OpenConfigd
PACKAGE_MAINTAINER:=ishi@coreswitch.io
PACKAGE_URL:=http://coreswitch.io/
PACKAGE_VENDOR:=coreswitch

OUT_DIR:=out
STAGING_DIR:=$(OUT_DIR)/stage
BUILD_DIR:=$(OUT_DIR)/build

.PHONY: package stage clean files

package: $(OUT_DIR)/timestamp.deb $(OUT_DIR)/timestamp.rpm

$(OUT_DIR)/timestamp.deb: stage
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
	touch $(OUT_DIR)/timestamp.deb

$(OUT_DIR)/timestamp.rpm: stage
	@echo "Creating package"
	@rm -f $(OUT_DIR)/$(PACKAGE_NAME)_$(PACKAGE_VERSION)*.rpm
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
	-t rpm .
	touch $(OUT_DIR)/timestamp.rpm

stage: files openconfigd cli

files: $(STAGING_DIR) $(STAGING_DIR)/etc $(STAGING_DIR)/etc/openconfigd
	rsync -avz ./files/ubuntu/ $(STAGING_DIR)/
	rsync -avz ../bash_completion.d $(STAGING_DIR)/etc
	rsync -avz ../yang $(STAGING_DIR)/etc/openconfigd

openconfigd: $(STAGING_DIR)/usr/bin
	go get github.com/coreswitch/openconfigd/openconfigd
	go get github.com/coreswitch/openconfigd/cli_command
	cp $(GOPATH)/bin/openconfigd $(STAGING_DIR)/usr/bin
	cp $(GOPATH)/bin/cli_command $(STAGING_DIR)/usr/bin

cli: $(BUILD_DIR) $(STAGING_DIR)/usr/bin
	(cd $(BUILD_DIR); ../../../cli/configure; make; cp cli ../../$(STAGING_DIR)/usr/bin)

$(STAGING_DIR):
	mkdir -p $(STAGING_DIR)

$(STAGING_DIR)/usr/bin:
	mkdir -p $(STAGING_DIR)/usr/bin

$(STAGING_DIR)/etc:
	mkdir -p $(STAGING_DIR)/etc

$(STAGING_DIR)/etc/openconfigd:
	mkdir -p $(STAGING_DIR)/etc/openconfigd

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(OUT_DIR)
