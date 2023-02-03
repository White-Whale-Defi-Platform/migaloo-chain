# Migaloo Testnet

This testnet will start with the node version `v1.0.0-rc0`.

**Genesis File**

To be finalized. Once finalize, it will be below:

```bash
# Get genesis file
curl -s https://raw.githubusercontent.com/White-Whale-Defi-Platform/migaloo-chain/main/networks/testnet/genesis.json > ~/.migalood/config/genesis.json

# Check sha256 (Should be TBD)
sha256sum ~/.migalood/config/genesis.json
```

**Node version**

```bash
# Get node version (should be v1.0.0-rc0)
migalood version --long | grep commit

# Get node long version (should be 78953dff50cf2f292a0f00eb6d7629531d86716d)
migalood version --long | grep commit
```

**Seed nodes**

```
TBD
```

**Persistent Peers**

```
TBD
```

## Setup

**Prerequisites:** Make sure to have [Golang >=1.19](https://golang.org/).

#### Go setup

You need to ensure your gopath configuration is correct. If the following **'make'** step does not work then you might have to add these lines to your .profile or .zshrc in the users home .migalood:

```sh
nano ~/.profile
```

```
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin
```

Source update .profile

```sh
source .profile
```

### Minimum hardware requirements

- 8-16GB RAM
- 100GB of disk space
- 2 cores

## Setup validator node

Below are the instructions to generate & submit your genesis transaction

### Generate genesis transaction (gentx)

Initialize the directories and create the local genesis file with the correct chain-id:

```bash
migalood init <moniker-name> --chain-id=test-chain-hU8r2x
```

Create a local key pair (skip this step if you already have a key):

```bash
migalood keys add KEY_NAME
```

Add your account to your local genesis file with a given amount and the key you just created. Use only `2000000uwhale`, other amounts will be ignored.

```bash
migalood add-genesis-account $(migalood keys show KEY_NAME -a) 2000000uwhale
```

Create the gentx, use only `1000000uwhale`:

```bash
migalood gentx KEY_NAME 1000000uwhale --chain-id=test-chain-hU8r2x
```

If all goes well, you will see a message similar to the following:

```bash
Genesis transaction written to "/home/user/.migalood/config/gentx/gentx-******.json"
```

### Submit genesis transaction

- Fork [the this repo](https://github.com/White-Whale-Defi-Platform/migaloo-chain) into your Github account

- Clone your repo using

```bash
git clone https://github.com/<your-github-username>/migaloo-chain
```

- Copy the generated gentx json file to `networks/testnet/gentx/`

```sh
cd migaloo-chain
cp ~/.migalood/config/gentx/gentx*.json ./networks/testnet/gentx/
```

- Commit and push to your repo
- Create a PR onto https://github.com/White-Whale-Defi-Platform/migaloo-chain
- Only PRs from individuals / groups with a history successfully running nodes will be accepted. This is to ensure the network successfully starts on time.

#### Running in production

If you have not installed cosmovisor, create a systemd file for your Juno service:

```sh
sudo nano /etc/systemd/system/migalood.service
```

Copy and paste the following and update `<YOUR_USERNAME>` and `<test-chain-hU8r2x>`:

```sh
Description=Juno daemon
After=network-online.target

[Service]
User=juno
ExecStart=/home/<YOUR_USERNAME>/go/bin/migalood start
Restart=on-failure
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
```

Enable and start the new service:

```sh
sudo systemctl enable migalood
sudo systemctl start migalood
```

Check status:

```sh
migalood status
```

Check logs:

```sh
journalctl -u migalood -f
```
