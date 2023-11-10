#!/usr/bin/env bash

set -eo pipefail

mkdir -p ./tmp-swagger-gen

cd proto
# Get the path of the cosmos-sdk repo from go/pkg/mod
cosmos_sdk_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/cosmos-sdk) || { echo "Error: Failed to find github.com/cosmos/cosmos-sdk"; exit 1; }
if [ -d "${cosmos_sdk_dir}/proto" ]; then
  proto_dirs=$(find ./migaloo "${cosmos_sdk_dir}/proto" -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
  for dir in $proto_dirs; do
    # generate swagger files (filter query files)
    query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))

    if [[ ! -z "$query_file" ]]; then
      buf generate --template buf.gen.swagger.yaml $query_file
      echo $query_file
    fi
  done
fi

cd ..
if [ -d "./client/docs" ]; then
  cd ./client/docs
  yarn install
  yarn combine
  yarn convert
  yarn build
  cd ../../
fi

# clean swagger files
rm -rf ./tmp-swagger-gen
