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

# remove existing daemon
rm -rf ~/.migaloo*

migalood config keyring-backend $KEYRING
migalood config chain-id $CHAINID

# if $KEY exists it should be deleted
echo "decorate bright ozone fork gallery riot bus exhaust worth way bone indoor calm squirrel merry zero scheme cotton until shop any excess stage laundry" | migalood keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --recover

migalood init $MONIKER --chain-id $CHAINID 

# Allocate genesis accounts (cosmos formatted addresses)
migalood add-genesis-account $KEY 100000000000000000000000000stake --keyring-backend $KEYRING

# Sign genesis transaction
migalood gentx $KEY 1000000000000000000000stake --keyring-backend $KEYRING --chain-id $CHAINID

# Collect genesis tx
migalood collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly 
migalood validate-genesis

if [[ $1 == "pending" ]]; then
  echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
migalood start --pruning=nothing  --minimum-gas-prices=0.0001stake --rpc.laddr tcp://0.0.0.0:26657
