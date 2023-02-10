KEY="mykey"
CHAINID="test-1"
MONIKER="localtestnet"
KEYALGO="secp256k1"
KEYRING="test"
LOGLEVEL="info"
# to trace evm
#TRACE="--trace"
TRACE=""

# validate dependencies are installed
command -v jq > /dev/null 2>&1 || { echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"; exit 1; }

# remove existing daemon

rm -rf ~/.migaloo*
migalood config keyring-backend $KEYRING
migalood config chain-id $CHAINID

# if $KEY exists it should be deleted
migalood keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO

# Set moniker and chain-id for Evmos (Moniker can be anything, chain-id must be an integer)
migalood init $MONIKER --chain-id $CHAINID 

migalood keys add key1 --pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AsiBD7b+DyvQ6Z71Rnijy0mmCIqi7Z8DsExNNQeVos69"}'
migalood keys add key2 --pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AjHEmOl57d5XuQuQQX738d9cgVoWW0YaVuvQntmw7nPa"}'
migalood keys add key3 --pubkey '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aw6j6+n4OyB0+NZnqedqNY14pa0mWFEfCktbao32BW5+"}'


# Allocate genesis accounts (cosmos formatted addresses)
migalood add-genesis-account $KEY 100000000000000000000000000stake --keyring-backend $KEYRING
migalood add-genesis-account key1 100000000000000000000000stake --keyring-backend $KEYRING
migalood add-genesis-account key2 100000000000000000000000stake --keyring-backend $KEYRING
migalood add-genesis-account key3 100000000000000000000000stake --keyring-backend $KEYRING

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
migalood start --pruning=nothing  --minimum-gas-prices=0.0001stake --rpc.laddr tcp://127.0.0.1:26650

