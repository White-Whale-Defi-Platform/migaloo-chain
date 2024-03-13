# Upgrade from Migaloo v4.1.2 to v4.1.3

> ## This is an important security update. IT IS CONSENSUS BREAKING, so please apply the fix only on height 5962800.

### Release Details
* https://github.com/White-Whale-Defi-Platform/migaloo-chain/releases/tag/v4.1.3
* Chain upgrade height : `5962800`. Exact upgrade time can be checked [here](https://ping.pub/whitewhale/block/5962800).


# To upgrade migaloo-chain

## Step 1: Alter systemd service configuration

We need to disable automatic restart of the node service. To do so please alter your `cosmovisor.service` file configuration and set appropriate lines to following values.

```
Restart=no 

Environment="DAEMON_ALLOW_DOWNLOAD_BINARIES=false"
Environment="DAEMON_RESTART_AFTER_UPGRADE=false"
```

After that you will need to run `sudo systemctl daemon-reload` to apply changes in the service configuration.

There is no need to restart the node yet; these changes will get applied during the node restart in the next step.

## Step 2: Restart migalood with a configured `halt-height`.

This upgrade requires `migalood` to have knowledge of the planned halt height. Please be aware that there is an extra step at the end to revert to `migalood`'s original configurations.

There are two mutually exclusive options for this stage:

### Option 1: Set the halt height by modifying `app.toml`

* Stop the `migalood` process.

* Edit the application configuration file at `~/.migalood/config/app.toml` so that `halt-height` reflects the upgrade plan:

```toml
# Note: Commitment of state will be attempted on the corresponding block.
halt-height = 5962800
```
* Start `migalood` process

* Wait for the upgrade height and confirm that the node has halted

### Option 2: Restart the `migalood` binary with command line flags

* Stop the `migalood` process.

* Do not modify `app.toml`. Restart the `migalood` process with the flag `--halt-height`:
```shell
migalood --halt-height 7818500
```

* Wait for the upgrade height and confirm that the node has halted

After performing these steps, the upgrade will proceed as usual using Cosmovisor.

# Setup Cosmovisor
## Create the updated migaloo binary of v4.1.3

### Go to migaloo-chain directory if present else clone the repository

```shell
   git clone https://github.com/White-Whale-Defi-Platform/migaloo-chain.git
```

### Follow these steps if migaloo-chain repo already present

```shell
   cd migaloo-chain
   git pull
   git fetch --tags
   git checkout v4.1.3
   make install
```

## Copy the new migaloo (v4.1.3) binary to cosmovisor current directory
```shell
   cp $GOPATH/bin/migalood ~/.migalood/cosmovisor/current/bin
```

## Restore service file settings

If you are using a service file, restore the previous `Restart` settings in your service file:
```
Restart=On-failure 
```
Reload the service control `sudo systemctl daemon-reload`.

# Revert `migalood` configurations

Depending on which path you chose for Step 1, either:

* Reset `halt-height = 0` option in the `app.toml` or
* Remove it from start parameters of the `migalood` binary and start node again