
<a name="v4.1.0"></a>
## [v4.1.0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v3.0.4...v4.1.0) (2024-02-16)

### Chore

* refactor format
* add changelog and update readme
* update go mod
* get upgrade handler
* change name v3 -> v4
* update make file
* fix update test genesis
* add third-party proto
* remove code not using
* update denom in the genesis files
* remove debug log
* add terra core 2.5.0 package

### Feat

* add checksum for wasmvm download
* make file to run test upgrade with multiple nodes
* base upgrade test
* setup cosmosvisor to run migalood
* init genesis with cosmovisor
* build linux migalood
* build cosmovisor migalood env linux images
* docker file to build v3.0.4
* docker setup to run migaloo
* fix ibc hooks
* add codec handler MsgUpdateTxFeeBurnPercentProposal
* add handler gov feeburn
* add cli
* add handler update
* add proposal types
* add proposal proto
* add feeburn to upgrade handler
* add feeburn module to app
* add testutil
* add feeburn clone from chihuahua
* add miss proto
* base test framework
* add script upgrade
* add upgrade v4.1.0
* update token factory
* alliance to 0.3.2

### Fix

* typo integration test
* add permission to account fee_collector
* upgrade name
* golint
* register ica controller
* ibc callback
* bump wasmvm to 1.5.1
* miss go.sum
* miss kill old migalood
* rollback go mod
* script unused
* using correct migaloo coin type
* miss subspace key table
* miss add chainId
* **test:** fix cycle import

### Perf

* add three more nodes
* not build if the image is already created
* reduce voting period
* set gas price to 0 stake for test

### Refactor

* add SetBech32 account
* update proto
* move v3 code to migalood-env
* golint ([#303](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/303))
* update script
* update name convention
* update name convention
* update UPGRADE_HEIGHT
* logic upgrade
* update wasm option
* update Makefile

### Test

* add test upgrade
* init with 50% feeburn
* ica integration test
* ibc callback with timeout
* ibc hooks
* add alliance test
* update test ante
* refactor app test
* add ante test
* adjust gas and fee doing tx token factory
* can init relayer


<a name="v3.0.4"></a>
## [v3.0.4](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v3.0.3...v3.0.4) (2024-01-10)

### Build

* **deps:** bump github.com/spf13/cobra from 1.6.1 to 1.8.0 ([#273](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/273))
* **deps:** bump github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v6 ([#266](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/266))
* **deps:** bump github.com/dvsekhvalnov/jose2go from 1.5.0 to 1.6.0 ([#308](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/308))
* **deps:** bump actions/upload-artifact from 3 to 4 ([#284](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/284))
* **deps:** bump golang.org/x/crypto from 0.9.0 to 0.17.0 ([#285](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/285))
* **deps:** bump github/codeql-action from 2 to 3 ([#283](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/283))
* **deps:** bump actions/setup-go from 4 to 5 ([#279](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/279))


<a name="v3.0.3"></a>
## [v3.0.3](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v3.0.2...v3.0.3) (2024-01-10)

### Build

* **deps:** bump github.com/prometheus/client_golang


<a name="v3.0.2"></a>
## [v3.0.2](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v3.0.1-hotfix...v3.0.2) (2023-11-06)


<a name="v3.0.1-hotfix"></a>
## [v3.0.1-hotfix](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.7-hotfix...v3.0.1-hotfix) (2023-10-09)

### Build

* **deps:** bump github.com/spf13/cast from 1.5.0 to 1.5.1
* **deps:** bump github.com/spf13/viper from 1.15.0 to 1.16.0
* **deps:** bump github.com/prometheus/client_golang
* **deps:** bump docker/setup-buildx-action from 2 to 3
* **deps:** bump docker/build-push-action from 4 to 5
* **deps:** bump docker/setup-qemu-action from 2 to 3
* **deps:** bump docker/login-action from 2 to 3
* **deps:** bump actions/checkout from 3 to 4


<a name="v2.2.7-hotfix"></a>
## [v2.2.7-hotfix](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v3.0.0...v2.2.7-hotfix) (2023-10-09)


<a name="v3.0.0"></a>
## [v3.0.0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.6...v3.0.0) (2023-09-26)

### Build

* **deps:** bump github.com/spf13/cast from 1.5.0 to 1.5.1
* **deps:** bump github.com/spf13/viper from 1.15.0 to 1.16.0
* **deps:** bump github.com/prometheus/client_golang
* **deps:** bump docker/setup-buildx-action from 2 to 3
* **deps:** bump docker/build-push-action from 4 to 5
* **deps:** bump docker/setup-qemu-action from 2 to 3
* **deps:** bump docker/login-action from 2 to 3
* **deps:** bump actions/checkout from 3 to 4


<a name="v2.2.6"></a>
## [v2.2.6](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.5...v2.2.6) (2023-07-19)

### Fix

* upgrade handler and store loader


<a name="v2.2.5"></a>
## [v2.2.5](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.4...v2.2.5) (2023-07-12)


<a name="v2.2.4"></a>
## [v2.2.4](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.3...v2.2.4) (2023-07-12)


<a name="v2.2.3"></a>
## [v2.2.3](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.2...v2.2.3) (2023-07-12)


<a name="v2.2.2"></a>
## [v2.2.2](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.1...v2.2.2) (2023-07-12)


<a name="v2.2.1"></a>
## [v2.2.1](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.2.0...v2.2.1) (2023-07-12)


<a name="v2.2.0"></a>
## [v2.2.0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.7...v2.2.0) (2023-07-12)


<a name="v2.0.7"></a>
## [v2.0.7](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.5...v2.0.7) (2023-07-06)


<a name="v2.0.5"></a>
## [v2.0.5](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.6...v2.0.5) (2023-06-09)


<a name="v2.0.6"></a>
## [v2.0.6](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.4...v2.0.6) (2023-06-09)


<a name="v2.0.4"></a>
## [v2.0.4](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.3...v2.0.4) (2023-06-09)


<a name="v2.0.3"></a>
## [v2.0.3](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.2...v2.0.3) (2023-05-26)


<a name="v2.0.2"></a>
## [v2.0.2](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.1...v2.0.2) (2023-04-18)


<a name="v2.0.1"></a>
## [v2.0.1](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.0...v2.0.1) (2023-03-28)

### Build

* **deps:** bump actions/checkout from 2 to 3


<a name="v2.0.0"></a>
## [v2.0.0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v2.0.0-rc0...v2.0.0) (2023-03-20)


<a name="v2.0.0-rc0"></a>
## [v2.0.0-rc0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v1.0.2...v2.0.0-rc0) (2023-03-19)

### Build

* **deps:** bump actions/setup-go from 3 to 4
* **deps:** bump cosmossdk.io/math from 1.0.0-beta.6 to 1.0.0-rc.0
* **deps:** bump codacy/codacy-analysis-cli-action from 4.2.0 to 4.3.0
* **deps:** bump github.com/cosmos/cosmos-sdk from 0.46.10 to 0.46.11


<a name="v1.0.2"></a>
## [v1.0.2](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v1.0.1...v1.0.2) (2023-03-19)


<a name="v1.0.1"></a>
## [v1.0.1](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v1.0.0...v1.0.1) (2023-03-19)

### Build

* **deps:** bump github.com/cosmos/cosmos-sdk from 0.46.9 to 0.46.10
* **deps:** bump github.com/hashicorp/go-getter from 1.6.1 to 1.7.0

### Feat

* **networks:** finalize genesis file

### Fix

* **networks:** typo in mainnet instructions
* **networks:** update mainnet instructions

### ShellCheck

*  modify path


<a name="v1.0.0"></a>
## [v1.0.0](https://github.com/White-Whale-Defi-Platform/migaloo-chain/compare/v1.0.0-rc0...v1.0.0) (2023-02-13)

### Build

* **deps:** bump github.com/spf13/viper from 1.14.0 to 1.15.0

### Feat

* **mainnet:** change app.toml to zero fees.
* **mainnet:** add mainnet instructions
* **mainnet:** add genesis-population.py
* **networks:** add genesis account
* **networks:** add genesis account
* **networks:** add community pool allocation
* **networks:** update script
* **networks:** add mainnet pre-genesis.json
* **networks:** update genesis.json

### Fix

* **networks:** regenerate pre-genesis.json
* **networks:** use more precise vesting timestamps
* **networks:** fix total supply
* **networks:** replace link to main pre-genesis.json
* **networks:** increase gentx allocation by 5 whale
* **networks:** fix bug in code snip


<a name="v1.0.0-rc0"></a>
## v1.0.0-rc0 (2023-02-03)

### Build

* **deps:** bump actions/dependency-review-action from 2 to 3 ([#27](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/27))
* **deps:** bump codacy/codacy-analysis-cli-action from 1.1.0 to 4.2.0 ([#28](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/28))
* **deps:** bump github.com/terra-money/alliance ([#29](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/29))
* **deps:** bump github.com/cosmos/interchain-accounts
* **deps:** bump docker/build-push-action from 3 to 4

### Chore

* go mod tidy

### Docs

* various docs improvements
* various docs improvements
* add Migaloo banner to readme
* add roadmap and contributing guide
* improve readme
* add security, code of conduct and contributing docs.

### Feat

* add alliance
* open source guidelines
* **.github:** Improve pull request and issues templates
* **.github:** Add pull request templates, issue templates and a
* **networks:** update genesis.json ([#9](https://github.com/White-Whale-Defi-Platform/migaloo-chain/issues/9))

