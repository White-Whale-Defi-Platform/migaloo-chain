#!/usr/bin/env bash

  

set -eo pipefail

  

mkdir -p ./tmp-swagger-gen

  

# move the vendor folder to a temp dir so that go list works properly

temp_dir="f29ea6aa861dc4b083e8e48f67cce"

if [ -d vendor ]; then

mv ./vendor ./$temp_dir

fi

  

go mod download github.com/terra-money/alliance

  

alliance_dir=$(go list -f '{{ .Dir }}' -m github.com/terra-money/alliance)

echo "alliance dir: $alliance_dir"

  
  

cd proto

proto_dirs=$(find $alliance_dir/proto ./migaloo -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)

  
  

for dir in $proto_dirs; do

# generate swagger files (filter query files)

query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))

echo "query file: $query_file"

if [[ ! -z "$query_file" ]]; then
  buf generate --template buf.gen.swagger.yaml $query_file
fi

done

cd ..

npx swagger-combine ./client/docs/config.json -o ./client/docs/swagger-ui/swagger.yaml -f yaml --continueOnConflictingPaths true --includeDefinitions true

# cat the content of the file
cat ./client/docs/swagger-ui/swagger.yaml

# clean swagger files
rm -rf ./tmp-swagger-gen