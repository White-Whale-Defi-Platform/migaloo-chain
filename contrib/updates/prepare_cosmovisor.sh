#!/bin/bash

# this bash will prepare cosmosvisor to the build folder so that it can perform upgrade
# this script is supposed to be run by Makefile

# These fields should be fetched automatically in the future
# Need to do more upgrade to see upgrade patterns
OLD_VERSION=v3.0.4
# this command will retrieve the folder with the largest number in format v<number>
SOFTWARE_UPGRADE_NAME=$(ls -d -- ./app/upgrades/v* | sort -Vr | head -n 1 | xargs basename)
BUILDDIR=$1
TESTNET_NVAL=$2
TESTNET_CHAINID=$3

# check if BUILDDIR is set
if [ -z "$BUILDDIR" ]; then
    echo "BUILDDIR is not set"
    exit 1
fi

# install old binary if not exist
if [ ! -f "_build/$OLD_VERSION.zip" ] &> /dev/null
then
    mkdir -p _build/old
    wget -c "https://github.com/White-Whale-Defi-Platform/migaloo-chain/archive/refs/tags/${OLD_VERSION}.zip" -O _build/${OLD_VERSION}.zip
    unzip _build/${OLD_VERSION}.zip -d _build
fi


if [ ! -f "$BUILDDIR/$OLD_VERSION.zip" ] &> /dev/null
then
    mkdir -p BUILDDIR/old
    docker build --platform linux/amd64 --no-cache --build-arg source=./_build/migaloo-chain-${OLD_VERSION:1}/ --tag migaloo/migalood.binary.old . 
    docker create --platform linux/amd64 --name old-temp migaloo/migalood.binary.old:latest
    docker cp old-temp:/usr/local/bin/migalood $BUILDDIR/old/
    docker rm old-temp
fi