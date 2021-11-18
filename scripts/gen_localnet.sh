#!/bin/sh

go build -o ./sisu cmd/sisud/main.go
if [ $? -eq 0 ]; then
    echo "Build succeeded"
else
    exit 1
fi

./sisu localnet --enable-tss true
rm -rf ~/.sisu
cp -rf ./output/node0/ ~/.sisu
rm -rf ./output

# Copy dheart.toml to its folder
mkdir -p ~/.sisu/dheart
cp ../dheart/dheart.toml ~/.sisu/dheart