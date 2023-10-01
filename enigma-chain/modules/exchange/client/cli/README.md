# Command examples

## Batch exchange modifications proposal

```sh
yes 12345678 | enigmad tx exchange batch-exchange-modifications-proposal --proposal="./enigma-chain/modules/exchange/client/cli/proposals/batchproposal.json" --deposit=10000000fury --from=genesis
```

## Batch community pool spend proposal

```sh
yes 12345678 | enigmad tx exchange batch-community-pool-spend-proposal --proposal="./enigma-chain/modules/exchange/client/cli/proposals/batchcommunitypoolspendproposal.json" --deposit=10000000fury --from=genesis
```
