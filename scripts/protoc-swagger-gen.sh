#!/usr/bin/env bash
set -eo pipefail

mkdir -p ./tmp-swagger-gen
# move the vendor folder to a temp dir so that go list works properly

temp_dir="f29ea6aa861dc4b083e8e48f67cce"

if [ -d vendor ]; then
  mv ./vendor ./$temp_dir
fi

# Define the version you want to clone
COSMOS_SDK_VERSION="v0.47.7"

# Define the directory where you want to clone the repo
COSMOS_SDK_DIR="./cosmos-sdk"

go mod download github.com/terra-money/alliance
go mod download github.com/cosmos/ibc-go/v7

# Clone the specific version of cosmos-sdk
if [ ! -d "$COSMOS_SDK_DIR" ]; then
  git clone --branch $COSMOS_SDK_VERSION --depth 1 https://github.com/cosmos/cosmos-sdk.git $COSMOS_SDK_DIR
else
  echo "Cosmos SDK directory already exists, skipping clone"
fi



alliance_dir=$(go list -f '{{ .Dir }}' -m github.com/terra-money/alliance)
ibc_dir=$(go list -f '{{ .Dir }}' -m github.com/cosmos/ibc-go/v7)

cd proto
proto_dirs=$(find ../$COSMOS_SDK_DIR/proto/cosmos $alliance_dir/proto $ibc_dir/proto ./migaloo -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))

  if [[ ! -z "$query_file" ]]; then
    echo "generated swagger file for $query_file"
    buf generate --template buf.gen.swagger.yaml $query_file
  fi
done

cd ..
npx swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# cat the content of the file
cat ./client/docs/swagger-ui/swagger.yaml

# clean swagger files
rm -rf ./tmp-swagger-gen