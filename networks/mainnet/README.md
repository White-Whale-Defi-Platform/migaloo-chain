# Migaloo Mainnet

This mainnet will start with the node version `b4fbc78`.

## Minimum hardware requirements

- 8-16GB RAM
- 100GB of disk space
- 2 cores

## Genesis Instruction

### Install node

```bash
git clone https://github.com/White-Whale-Defi-Platform/migaloo-chain
cd migaloo-chain
git checkout b4fbc78
make install
```

### Check Node version

```bash
# Get node version (should be b4fbc78)
migalood version

# Get node long version (should be ???)
migalood version --long | grep commit
```

### Initialize Chain

```bash
migalood init MONIKER --chain-id=migaloo-1
```

Set minimum gas price to 0.

```bash
sed -i 's/minimum-gas-prices = ".*"/minimum-gas-prices = "0uwhale"/' $HOME/.migalood/config/app.toml
```

### Download pre-genesis

```bash
curl -s https://raw.githubusercontent.com/White-Whale-Defi-Platform/migaloo-chain/release/v1.0.x/networks/mainnet/pre-genesis.json > ~/.migalood/config/genesis.json
```

## Create gentx

Create wallet

```bash
migalood keys add KEY_NAME
```

Fund yourself `25000000uwhale`

```bash
migalood add-genesis-account $(migalood keys show KEY_NAME -a) 25000000uwhale
```

Use half (`10000000uwhale`) for self-delegation

```bash
migalood gentx KEY_NAME 10000000uwhale --chain-id=migaloo-1
```

If all goes well, you will see a message similar to the following:

```bash
Genesis transaction written to "/home/user/.migalood/config/gentx/gentx-******.json"
```

### Submit genesis transaction

- Fork this repo
- Copy the generated gentx json file to `networks/mainnet/gentx/`
- Commit and push to your repo
- Create a PR on this repo
