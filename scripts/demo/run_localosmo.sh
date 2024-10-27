#!/bin/sh

rm -rf $HOME/.osmosisd

$HOME/go/bin/osmosisd init mac --chain-id test-osmo

cp $HOME/dev/cosmos/interchain-security/scripts/demo/osmosis/genesis.json $HOME/.osmosisd/config/genesis.json
cp $HOME/dev/cosmos/interchain-security/scripts/demo/osmosis/config.toml $HOME/.osmosisd/config/config.toml
cp $HOME/dev/cosmos/interchain-security/scripts/demo/osmosis/app.toml $HOME/.osmosisd/config/app.toml

$HOME/go/bin/osmosisd add-genesis-account key1 10000000000000000uosmo
$HOME/go/bin/osmosisd add-genesis-account relayer 10000000000000000uosmo
$HOME/go/bin/osmosisd add-genesis-account osmo1dhat6w6256krheukuzwp8m4z5d6za9zx2jachsjljpmq7t39hzlqrzx8fc 10000000000usomo
$HOME/go/bin/osmosisd gentx key1 100000000000000uosmo --chain-id test-osmo
$HOME/go/bin/osmosisd collect-gentxs
$HOME/go/bin/osmosisd start
