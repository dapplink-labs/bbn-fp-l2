BUILDDIR ?= $(CURDIR)/build
TOOLS_DIR := tools

BABYLON_PKG := github.com/babylonlabs-io/babylon/cmd/babylond

MANTA_STAKING_ARTIFACT := ./manta-staking-contracts/out/MantaStakingMiddleware.sol/MantaStakingMiddleware.json

GO_BIN := ${GOPATH}/bin
BTCD_BIN := $(GO_BIN)/btcd

DOCKER := $(shell which docker)
CUR_DIR := $(shell pwd)
MOCKS_DIR=$(CUR_DIR)/testutil/mocks
MOCKGEN_REPO=github.com/golang/mock/mockgen
MOCKGEN_VERSION=v1.6.0
MOCKGEN_CMD=go run ${MOCKGEN_REPO}@${MOCKGEN_VERSION}

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')

ldflags := $(LDFLAGS) -X github.com/Manta-Network/manta-fp/version.version=$(VERSION)
build_tags := $(BUILD_TAGS)
build_args := $(BUILD_ARGS)

PACKAGES_E2E=$(shell go list ./... | grep '/itest')

ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static" -v
endif

ifeq ($(VERBOSE),true)
	build_args += -v
endif

BUILD_TARGETS := build install
BUILD_FLAGS := --tags "$(build_tags)" --ldflags '$(ldflags)'

# Update changelog vars
ifneq (,$(SINCE_TAG))
	sinceTag := --since-tag $(SINCE_TAG)
endif
ifneq (,$(UPCOMING_TAG))
	upcomingTag := --future-release $(UPCOMING_TAG)
endif

all: build install

build: BUILD_ARGS := $(build_args) -o $(BUILDDIR)

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	CGO_CFLAGS="-O -D__BLST_PORTABLE__" go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-docker:
	$(DOCKER) build --tag Manta-Network/manta-fp -f Dockerfile \
		$(shell git rev-parse --show-toplevel)

.PHONY: build build-docker

.PHONY: test
test:
	go test ./...

test-e2e:
	go test -mod=readonly -failfast -timeout=25m -v $(PACKAGES_E2E) -count=1 --tags=e2e

###############################################################################
###                                Protobuf                                 ###
###############################################################################

proto-all: proto-gen

proto-gen:
	make -C eotsmanager proto-gen
	make -C finality-provider proto-gen

binding:
	$(eval temp := $(shell mktemp))

	cat $(MANTA_STAKING_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(MANTA_STAKING_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/manta_staking_middleware.go \
		--type MantaStakingMiddleware \
		--bin $(temp)

		rm $(temp)

.PHONY: proto-gen
