GOBIN=go
GORUN=$(GOBIN) run
GOBUILD=$(GOBIN) build
GOTEST=$(GOBIN) test

BUILD_PATH=$(shell pwd)/build

BENCHTIME=1s
BENCHTIMEOUT=10m

all: collector-build

clean:
	rm -rf $(BUILD_PATH)/*

collector-build:
	mkdir -p $(BUILD_PATH)/collector
	$(GOBUILD) -o $(BUILD_PATH)/collector ./cmd/server/server.go

codacy-coverage-push:
	$(GOTEST) -coverprofile=coverage.out ./...
	bash scripts/get.sh report --force-coverage-parser go -r ./coverage.out

.PHONY: gen-src-archive
gen-src-archive:
	bash scripts/gen_src_archive.sh
