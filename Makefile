###############################################################################
# Licensed Materials - Property of IBM Copyright IBM Corporation 2017, 2019. All Rights Reserved.
# U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP
# Schedule Contract with IBM Corp.
#
# Contributors:
#  IBM Corporation - initial API and implementation
###############################################################################
#
# WARNING: DO NOT MODIFY. Changes may be overwritten in future updates.
#
# The following build goals are designed to be generic for any docker image.
# This Makefile is designed to be included in other Makefiles.
# You must ensure that Make variables are defined for IMAGE_REPO and IMAGE_NAME.
#
# If you are using a Bluemix image registry, you must also define BLUEMIX_API_KEY,
# BLUEMIX_ORG, and BLUEMIX_SPACE
###############################################################################

TAG_VERSION ?= `cat VERSION`+$(GIT_COMMIT)

.DEFAULT_GOAL := all

BEFORE_SCRIPT := $(shell ./build-tools/before-make-script.sh)

.PHONY: all-checks
all-checks: copyright-check

.PHONY: dep-install
dep-install::
	
	mkdir -p $(GOPATH)/bin
	dep version; \
	if [ $$? -ne 0 ]; then \
		curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh; \
	fi

.PHONY: pre-req
pre-req:: dep-install
	dep ensure -v


.PHONY: go-test 
go-test:
	@echo "Start Integration test"; 
	#go clean -testcache
	go test -p 1 -v -coverprofile=c.out github.com/IBM/s3mem-go/s3mem/... || exit 1
	go tool cover -func=c.out
	
.PHONY: copyright-check
copyright-check:
	./build-tools/copyright-check.sh

.PHONY: tag
tag::
	$(eval GIT_COMMIT := $(shell git rev-parse --short HEAD))
	@echo "TAG_VERSION:$(TAG_VERSION)"

.PHONY: all
all:: clean pre-req copyright-check go-test

.PHONY: clean
clean::
	rm -rf api/testFile
	