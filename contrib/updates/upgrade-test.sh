#!/bin/bash

# should make this auto fetch upgrade name from app upgrades once many upgrades have been done
# this command will retrieve the folder with the largest number in format v<number>
SOFTWARE_UPGRADE_NAME=$(ls -d -- ./app/upgrades/v* | sort -Vr | head -n 1 | xargs basename)
NODE1_HOME=node1/migalood
SELECTED_CONTAINER=migaloodnode1
BINARY_OLD="docker exec $SELECTED_CONTAINER ./old/migalood"
TESTNET_NVAL=${1:-3}

# sleep to wait for localnet to come up
echo "Wait for localnet to come up"
sleep 5

# 20 block from now
$BINARY_OLD status --home $NODE1_HOME 
STATUS_INFO=($($BINARY_OLD status --home $NODE1_HOME | jq -r '.NodeInfo.network,.SyncInfo.latest_block_height'))
echo "Current status info: $STATUS_INFO"
CHAIN_ID=${STATUS_INFO[0]}
UPGRADE_HEIGHT=$((STATUS_INFO[1] + 40))
echo "Upgrade should happens at: $UPGRADE_HEIGHT"


docker exec $SELECTED_CONTAINER tar -cf ./migalood.tar -C . migalood
SUM=$(docker exec $SELECTED_CONTAINER sha256sum ./migalood.tar | cut -d ' ' -f1)
DOCKER_BASE_PATH=$(docker exec $SELECTED_CONTAINER pwd)
echo $SUM
UPGRADE_INFO=$(jq -n '
{
    "binaries": {
        "linux/amd64": "file://'$DOCKER_BASE_PATH'/migalood.tar?checksum=sha256:'"$SUM"'",
    }
}')

echo $UPGRADE_INFO

echo "Submitting software upgrade proposal..."
$BINARY_OLD tx gov submit-legacy-proposal software-upgrade "$SOFTWARE_UPGRADE_NAME" --upgrade-height $UPGRADE_HEIGHT --upgrade-info "$UPGRADE_INFO" --title "upgrade" --description "upgrade"  --from node1 --keyring-backend test --chain-id $CHAIN_ID --home $NODE1_HOME -y > /dev/null 2>&1

sleep 5

echo "Depositing to software upgrade proposal..."
$BINARY_OLD tx gov deposit 1 "20000000stake" --from node1 --keyring-backend test --chain-id $CHAIN_ID --home $NODE1_HOME -y > /dev/null 2>&1

sleep 5

# loop from 0 to TESTNET_NVAL
for (( i=0; i<$TESTNET_NVAL; i++ )); do
    # check if docker for node i is running
    if [[ $(docker ps -a | grep migaloodnode$i | wc -l) -eq 1 ]]; then
        $BINARY_OLD tx gov vote 1 yes --from node$i --keyring-backend test --chain-id $CHAIN_ID --home "node$i/migalood" -y > /dev/null 2>&1
        echo -e "---> Node $i voted yes"
        sleep 1
    fi
done

# keep track of block_height
NIL_BLOCK=0
LAST_BLOCK=0
SAME_BLOCK=0
while true; do 
    BLOCK_HEIGHT=$($BINARY_OLD status --home $NODE1_HOME | jq '.SyncInfo.latest_block_height' -r)
    # if BLOCK_HEIGHT is empty
    if [[ -z $BLOCK_HEIGHT ]]; then
        # if 5 nil blocks in a row, exit
        if [[ $NIL_BLOCK -ge 5 ]]; then
            echo "ERROR: 5 nil blocks in a row"
            break
        fi
        NIL_BLOCK=$((NIL_BLOCK + 1))
    fi

    # if block height is not nil
    # if block height is same as last block height
    if [[ $BLOCK_HEIGHT -eq $LAST_BLOCK ]]; then
        # if 5 same blocks in a row, exit
        if [[ $SAME_BLOCK -ge 5 ]]; then
            echo "ERROR: 5 same blocks in a row"
            break
        fi
        SAME_BLOCK=$((SAME_BLOCK + 1))
    else
        # update LAST_BLOCK and reset SAME_BLOCK
        LAST_BLOCK=$BLOCK_HEIGHT
        SAME_BLOCK=0
    fi

    if [[ $BLOCK_HEIGHT -ge $UPGRADE_HEIGHT ]]; then
        # assuming running only 1 migalood
        echo "UPGRADE REACHED, CONTINUING NEW CHAIN"
        break
    else
        $BINARY_OLD q gov proposal 1 --output=json --home $NODE1_HOME | jq ".status"
        echo "BLOCK_HEIGHT = $BLOCK_HEIGHT"
        sleep 10
    fi
done

if [[ $SAME_BLOCK -ge 5 ]]; then
    docker logs migaloodnode0
    exit 1
fi

sleep 40

# check all nodes are online after upgrade
for (( i=0; i<$TESTNET_NVAL; i++ )); do
    if [[ $(docker ps -a | grep migaloodnode$i | wc -l) -eq 1 ]]; then
        docker exec migaloodnode$i ./migalood status --home "node$i/migalood"
        if [[ "${PIPESTATUS[0]}" != "0" ]]; then
            echo "node$i is not online"
            docker logs migaloodnode$i
            exit 1
        fi
    else
        echo "migaloodnode$i is not running"
        docker logs migaloodnode$i
        exit 1
    fi
done