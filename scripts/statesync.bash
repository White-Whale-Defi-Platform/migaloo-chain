#!/bin/bash
# microtick and bitcanna contributed significantly here.
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
wget -O ~/.migalood/config/genesis.json https://raw.githubusercontent.com/White-Whale-Defi-Platform/migaloo-chain/main/networks/mainnet/genesis.json

# Set minimum gas price.
sed -i'' 's/minimum-gas-prices = ""/minimum-gas-prices = "0.0025uwhale"/' $HOME/.migalood/config/app.toml

# Get "trust_hash" and "trust_height".
INTERVAL=1000
LATEST_HEIGHT=$(curl -s https://https://rpc-whitewhale.goldenratiostaking.net/block | jq -r .result.block.header.height)
BLOCK_HEIGHT=$(($LATEST_HEIGHT-$INTERVAL)) 
TRUST_HASH=$(curl -s "https://https://rpc-whitewhale.goldenratiostaking.net/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

# Print out block and transaction hash from which to sync state.
echo "trust_height: $BLOCK_HEIGHT"
echo "trust_hash: $TRUST_HASH"

# Export state sync variables.
export MIGALOOD_STATESYNC_ENABLE=true
export MIGALOOD_P2P_MAX_NUM_OUTBOUND_PEERS=200
export MIGALOOD_STATESYNC_RPC_SERVERS="https://https://whitewhale-rpc.lavenderfive.com:443,https://rpc-whitewhale.goldenratiostaking.net:443"
export MIGALOOD_STATESYNC_TRUST_HEIGHT=$BLOCK_HEIGHT
export MIGALOOD_STATESYNC_TRUST_HASH=$TRUST_HASH

# Fetch and set list of seeds from chain registry.
export MIGALOOD_P2P_SEEDS=$(curl -s https://raw.githubusercontent.com/cosmos/chain-registry/master/cosmoshub/chain.json | jq -r '[foreach .peers.seeds[] as $item (""; "\($item.id)@\($item.address)")] | join(",")')

# Start chain.
migalood start --x-crisis-skip-assert-invariants