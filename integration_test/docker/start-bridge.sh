#!/usr/bin/env bash

apt-get update
apt-get install -y curl jq

./celestia bridge --node.store /bridge init
/wait-for-it.sh -t 90 192.167.10.10:26657/header -- \
  curl -s http://192.167.10.10:26657/block?height=1 | jq '.result.block_id.hash' | tr -d '"' > genesis.hash

export CELESTIA_CUSTOM=ephemeral:`cat genesis.hash`
echo $CELESTIA_CUSTOM
./celestia bridge --node.store /bridge --rpc.port 26658 \
  --core.grpc 192.167.10.10:9090 \
  --core.remote tcp://192.167.10.10:26657 \
  --keyring.accname node \
  start
