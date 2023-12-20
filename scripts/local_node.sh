#!/bin/bash

KEY="mykey"
CHAINID="test-1"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

HOMEDIR="$HOME/.migaloo-tmp"

CONFIG=$HOMEDIR/config/config.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

# used to exit on first error (any non-zero exit code)
set -e

rm -rf "$HOMEDIR"

migalood config keyring-backend $KEYRING --home $HOMEDIR
migalood config chain-id $CHAINID --home $HOMEDIR

# if $KEY exists it should be deleted
migalood keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home $HOMEDIR

migalood init $MONIKER --chain-id $CHAINID --home $HOMEDIR

# Change parameter token denominations to uwhale
jq '.app_state["staking"]["params"]["bond_denom"]="uwhale"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["mint"]["params"]["mint_denom"]="uwhale"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["crisis"]["constant_fee"]["denom"]="uwhale"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="uwhale"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["tokenfactory"]["params"]["denom_creation_fee"][0]["denom"]="uwhale"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"


# Set gas limit in genesis
jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

if [[ $1 == "pending" ]]; then
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
        sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
        sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
        sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
        sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
        sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
        sed -i '' 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
        sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
    else
        sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
        sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
        sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
        sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
        sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
        sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
        sed -i 's/timeout_commit = "5s"/timeout_commit = "150s"/g' "$CONFIG"
        sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
    fi
fi

if [[ "$OSTYPE" == "darwin"* ]]; then
    sed -i '' 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG"
else
    sed -i 's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' "$CONFIG"
fi

# Allocate genesis accounts (cosmos formatted addresses)
migalood add-genesis-account $KEY 100000000000000000000000000uwhale --keyring-backend $KEYRING --home $HOMEDIR

# Sign genesis transaction
migalood gentx $KEY 1000000000000000000000uwhale --keyring-backend $KEYRING --chain-id $CHAINID --home $HOMEDIR

# Collect genesis tx
migalood collect-gentxs --home $HOMEDIR

# Run this to ensure everything worked and that the genesis file is setup correctly
migalood validate-genesis --home $HOMEDIR

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# update request max size so that we can upload the light client
# '' -e is a must have params on mac, if use linux please delete before run
sed -i'' -e 's/max_body_bytes = /max_body_bytes = 1/g' "$CONFIG"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
migalood start --pruning=nothing  --minimum-gas-prices=0.0001uwhale --rpc.laddr tcp://0.0.0.0:26657 --home $HOMEDIR