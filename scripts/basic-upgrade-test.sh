#!/bin/bash

PASSPHRASE="12345678"
TX_OPTS="--from=genesis  --chain-id=enigma-1 --gas-prices 500000000fury --broadcast-mode=block --yes"

# calculate halt height
CUR_HEIGHT=$(curl -sS localhost:26657/block | jq .result.block.header.height | tr -d '"')
HALT_HEIGHT=$(($CUR_HEIGHT + 5))

fetch_proposal_id() {
  current_proposal_id=$(curl 'http://localhost:10337/cosmos/gov/v1beta1/proposals?proposal_status=0&pagination.limit=1&pagination.reverse=true' | jq -r '.proposals[].proposal_id')
  proposal=$((current_proposal_id))
}

vote() {
        PROPOSAL_ID=$1
        echo $PROPOSAL_ID
        yes $PASSPHRASE | enigmad tx gov vote $PROPOSAL_ID yes $TX_OPTS
}


yes $PASSPHRASE | enigmad tx gov submit-proposal software-upgrade v1.11 \
 --title "Enigma Protocol 1.11 Dry Run" \
 --description "Enigma Upgrade 1.11" \
 --upgrade-height $HALT_HEIGHT \
 --deposit 500000000000000000000fury $TX_OPTS

fetch_proposal_id
vote $proposal
