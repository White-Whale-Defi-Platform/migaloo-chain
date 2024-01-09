#!/usr/bin/make -f

BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
COMMIT := $(shell git log -1 --format='%H')

APP_DIR = ./app
BINDIR ?= ~/go/bin
RUNSIM  = $(BINDIR)/runsim
BINARY ?= migalood

ifeq (,$(VERSION))
  VERSION := $(shell git describe --tags)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/cometbft/cometbft | sed 's:.* ::')
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
BUILDDIR ?= $(CURDIR)/build
HTTPS_GIT := https://github.com/White-Whale-Defi-Platform/migaloo-chain.git

export GO111MODULE = on

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  build_tags += gcc cleveldb
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=migaloo \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=migalood \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/cometbft/cometbft/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
  ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(COSMOS_BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif


all: install

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/migalood

build:
	go build $(BUILD_FLAGS) -o bin/migalood ./cmd/migalood

docker-build-debug:
	@DOCKER_BUILDKIT=1 docker build -t migaloo:debug -f Dockerfile .

runsim: $(RUNSIM)
$(RUNSIM):
	@echo "Installing runsim..."
	@go install github.com/cosmos/tools/cmd/runsim@v1.0.0

test-sim-import-export: runsim
	@echo "Running application import/export simulation. This may take several minutes..."
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(APP_DIR) 50 5 TestAppImportExport

test-sim-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.migalood/config/genesis.json will be used."
	@$(BINDIR)/runsim -Genesis=${HOME}/.migalood/config/genesis.json -SimAppPkg=$(APP_DIR) 400 5 TestFullAppSimulation

test-sim-multi-seed-long: runsim
	@echo "Running long multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(APP_DIR) 500 50 TestFullAppSimulation

test-sim-multi-seed-short: runsim
	@echo "Running short multi-seed application simulation. This may take awhile!"
	@$(BINDIR)/runsim -Jobs=4 -SimAppPkg=$(APP_DIR) 50 10 TestFullAppSimulation

test-sim-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.migalood/config/genesis.json will be used."
	@go test $(TEST_FLAGS) -mod=readonly $(SIMAPP) -run TestFullAppSimulation \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

###############################################################################
###                             Interchain test                             ###
###############################################################################

# Executes start chain tests via interchaintest
ictest-start-cosmos:
	cd tests/interchaintest && go test -race -v -run TestStartMigaloo .

ictest-ibc:
	cd tests/interchaintest && go test -race -v -run TestMigalooGaiaIBCTransfer .

ictest-ibc-hooks:
	cd tests/interchaintest && go test -race -v -run TestIBCHooks .

# Executes all tests via interchaintest after compling a local image as migaloo:local
ictest-all: ictest-start-cosmos ictest-ibc


###############################################################################
###                        Integration Tests                                ###
###############################################################################

#./scripts/tests/relayer/interchain-acc-config/rly-init.sh
	
init-test-framework: clean-testing-data install
	@echo "Initializing both blockchains..."
	./scripts/tests/init-test-framework.sh
	./scripts/tests/relayer/interchain-acc-config/rly-init.sh

test-tokenfactory: 
	@echo "Testing tokenfactory..."
	./scripts/tests/tokenfactory/tokenfactory.sh

test-alliance: 
	@echo "Testing alliance..."
	./scripts/tests/alliance/delegate.sh

test-ica:
	@echo "Testing ica..."
	./scripts/tests/ica/delegate.sh

test-ibc-hooks:
	@echo "Testing ibc-hooks..."
	./scripts/tests/ibc-hooks/increment.sh

clean-testing-data:
	@echo "Killing migallod and removing previous data"
	-@pkill $(BINARY) 2>/dev/null
	-@pkill rly 2>/dev/null
	-@pkill migalood_new 2>/dev/null
	-@pkill migalood_old 2>/dev/null
	-@rm -rf ./data

	

.PHONY: ictest-start-cosmos ictest-all ictest-ibc-hooks ictest-ibc
###############################################################################
###                                  Proto                                  ###
###############################################################################

proto-all: proto-format proto-gen

proto:
	@echo
	@echo "=========== Generate Message ============"
	@echo
	./scripts/protocgen.sh
	@echo
	@echo "=========== Generate Complete ============"
	@echo

test:
	@go test -v ./x/...

docs:
	@echo
	@echo "=========== Generate Message ============"
	@echo
	./scripts/generate-docs.sh

	statik -src=client/docs/static -dest=client/docs -f -m
	@if [ -n "$(git status --porcelain)" ]; then \
        echo "\033[91mSwagger docs are out of sync!!!\033[0m";\
        exit 1;\
    else \
        echo "\033[92mSwagger docs are in sync\033[0m";\
    fi
	@echo
	@echo "=========== Generate Complete ============"
	@echo

###############################################################################
###                                Protobuf                                 ###
###############################################################################

containerProtoVer=0.13.0
containerProtoImage=ghcr.io/cosmos/proto-builder:$(containerProtoVer)

proto-gen:
	@echo "Generating Protobuf files"
	@$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) \
		sh ./scripts/protocgen.sh;

proto-format:
	@echo "Formatting Protobuf files"
	@if docker ps -a --format '{{.Names}}' | grep -Eq "^${containerProtoFmt}$$"; then docker start -a $(containerProtoFmt); else docker run --name $(containerProtoFmt) -v $(CURDIR):/workspace --workdir /workspace tendermintdev/docker-build-proto \
		find ./ -not -path "./third_party/*" -name "*.proto" -exec clang-format -i {} \; ; fi
