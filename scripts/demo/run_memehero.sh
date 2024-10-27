#!/bin/sh

rm -rf $HOME/.meme-p

$HOME/go/bin/interchain-security-pd init mac --chain-id test-memep

cp $HOME/dev/cosmos/interchain-security/scripts/demo/memehero/config.toml $HOME/.meme-p/config/config.toml

$HOME/go/bin/interchain-security-pd genesis add-genesis-account key1 10000000000000000stake,10000000000000000uatom
$HOME/go/bin/interchain-security-pd genesis add-genesis-account relayer 100000000000000stake
$HOME/go/bin/interchain-security-pd genesis gentx key1 100000000000000stake --chain-id test-memep
$HOME/go/bin/interchain-security-pd genesis collect-gentxs
$HOME/go/bin/interchain-security-pd start