#!/bin/sh

hermes create client --host-chain test-memep --reference-chain test-osmo
hermes create client --host-chain test-osmo --reference-chain test-memep

hermes create connection --a-chain test-memep --a-client 07-tendermint-0 --b-client 07-tendermint-0
hermes create channel --a-chain test-memep --a-connection connection-0 --a-port transfer --b-port transfer
hermes startz