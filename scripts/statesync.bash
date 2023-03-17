#!/bin/bash
# microtick and bitcanna contributed significantly here.
# invoke this script in your migaloo-chain folder like this:
# bash scripts/statesync.bash
set -uxe

# Set Golang environment variables.
export GOPATH=~/go
export PATH=$PATH:~/go/bin

# Install Migaloo.
go install ./...

# NOTE: ABOVE YOU CAN USE ALTERNATIVE DATABASES, HERE ARE THE EXACT COMMANDS
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=rocksdb' -tags rocksdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=badgerdb' -tags badgerdb ./...
# go install -ldflags '-w -s -X github.com/cosmos/cosmos-sdk/types.DBBackend=boltdb' -tags boltdb ./...
# Tendermint team is currently focusing efforts on badgerdb.

# Initialize chain.
migalood init test 

# get genesis
wget -O ~/.migalood/config/genesis.json https://github.com/White-Whale-Defi-Platform/migaloo-chain/raw/release/v2.0.x/networks/mainnet/genesis.json


# Set minimum gas price.
sed -i '' 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0025uwhale"/' "$HOME/.migalood/config/app.toml"

# Get "trust_hash" and "trust_height".
INTERVAL=1000
LATEST_HEIGHT=$(curl -s https://rpc-whitewhale-h93nh9ykmqnzbrez-ie.internalendpoints.notional.ventures/block | jq -r .result.block.header.height)
BLOCK_HEIGHT=$((LATEST_HEIGHT-INTERVAL)) 
TRUST_HASH=$(curl -s "https://rpc-whitewhale-h93nh9ykmqnzbrez-ie.internalendpoints.notional.ventures/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables.
export MIGALOOD_STATESYNC_ENABLE=true
export MIGALOOD_P2P_MAX_NUM_OUTBOUND_PEERS=200
# replace the url below with a working one
export MIGALOOD_STATESYNC_RPC_SERVERS="http://65.108.5.173:2481"
export MIGALOOD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export MIGALOOD_STATESYNC_TRUST_HASH=$TRUST_HASH

# Fetch and set list of seeds from chain registry.
MIGALOOD_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/migaloo/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')
export MIGALOOD_P2P_SEEDS


# Start chain.
migalood start --x-crisis-skip-assert-invariants --minimum-gas-prices 0.00001uwhale