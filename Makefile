
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################

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
		go get -u github.com/golang/dep/cmd/dep; \
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

.PHONY: all
all:: clean pre-req copyright-check go-test

.PHONY: clean
clean::
	rm -rf api/testFile
	