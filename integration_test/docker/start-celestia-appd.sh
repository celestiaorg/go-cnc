#!/usr/bin/env bash

chain_id=ephemeral
token=celestia
node_name=node
key_type=test
celestia_appd=/celestia-app/celestia-appd

rm -rf $node_name/config/genesis.json $node_name/config/gentx
pwd
$celestia_appd --home $node_name init $node_name --chain-id $chain_id
$celestia_appd --home $node_name keys add $node_name --keyring-backend $key_type
addr=`$celestia_appd --home $node_name keys show $node_name --address --keyring-backend test`
echo $addr
$celestia_appd --home $node_name add-genesis-account $addr 800000000000$token

apk add moreutils # required for sponge
/fix-genesis.sh $node_name/config/genesis.json $token


$celestia_appd --home $node_name gentx $node_name 5000000000$token --keyring-backend=$key_type --chain-id $chain_id
$celestia_appd --home $node_name collect-gentxs

sed -i $node_name/config/config.toml -e 's/"full"/"validator"/'
$celestia_appd --home $node_name --rpc.laddr tcp://0.0.0.0:26657 start
