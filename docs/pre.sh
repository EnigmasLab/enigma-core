#!/usr/bin/env bash

#rm -rf ./docs/develop/modules/Core
rm -rf ./docs/develop/modules/Enigma
#mkdir ./docs/develop/modules/Core
mkdir ./docs/develop/modules/Enigma


# Get Enigma modules specs 

for D in ../enigma-chain/modules/*; do
  if [ -d "${D}" ]; then
    #rm -rf "modules/Enigma/$(echo $D | awk -F/ '{print $NF}')"
    mkdir -p "docs/develop/modules/Enigma/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
  fi
done

## Include the specs from latest Cosmos SDK
#git clone --depth 1 https://github.com/cosmos/cosmos-sdk.git
#
#for D in ./cosmos-sdk/x/*; do
#  if [ -d "${D}" ]; then
#    #rm -rf "modules/Enigma/$(echo $D | awk -F/ '{print $NF}')"
#    mkdir -p "docs/develop/modules/Core/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/README.md "$_"
#  fi
#done
#
#rm -rf cosmos-sdk

# Include the specs from Cosmos SDK v0.45.11
#git clone --depth 1 --branch v0.45.11 https://github.com/cosmos/cosmos-sdk.git
#
#for D in ./cosmos-sdk/x/*; do
#  if [ -d "${D}" ]; then
#    #rm -rf "modules/Enigma/$(echo $D | awk -F/ '{print $NF}')"
#    mkdir -p "docs/develop/modules/Core/$(echo $D | awk -F/ '{print $NF}')" && cp -r $D/spec/* "$_"
#    perl -i -pe 's/simd/enigmad/g' docs/develop/modules/Core/$(echo $D | awk -F/ '{print $NF}')/*
#  fi
#done

#rm -rf cosmos-sdk