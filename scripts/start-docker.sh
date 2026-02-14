#!/bin/bash

KEY="dev0"
CHAINID="qie_1990-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t qie-datadir.XXXXX)

echo "create and add new keys"
./qied keys add $KEY --home "$DATA_DIR" --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Qie with moniker=$MONIKER and chain-id=$CHAINID"
./qied init $MONIKER --chain-id "$CHAINID" --home "$DATA_DIR"
echo "prepare genesis: Allocate genesis accounts"
./qied add-genesis-account \
	"$(./qied keys show "$KEY" -a --home "$DATA_DIR" --keyring-backend test)" 1000000000000000000aqie,1000000000000000000stake \
	--home "$DATA_DIR" --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./qied gentx "$KEY" 1000000000000000000stake --keyring-backend test --home "$DATA_DIR" --keyring-backend test --chain-id "$CHAINID"
echo "prepare genesis: Collect genesis tx"
./qied collect-gentxs --home "$DATA_DIR"
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./qied validate-genesis --home "$DATA_DIR"

echo "starting qie node in background ..."
./qied start --pruning=nothing --rpc.unsafe \
	--keyring-backend test --home "$DATA_DIR" \
	>"$DATA_DIR"/node.log 2>&1 &
disown

echo "started qie node"
tail -f /dev/null
