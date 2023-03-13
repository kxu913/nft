## Blockchain base on Flow
*基于Flow的框架，实现一个用户在区块链上点赞功能*

### check account
https://flow-view-source.com/testnet/account/0xbbbe32d615d7c84b/keys
### check tx
https://flow-view-source.com/testnet/tx/e522d18b89cc43dec3690735f819117cac004473a61d7f57f9808f2466e42053

### basic flow cli command
- Execute script
```flow scripts execute ./cadence/scripts/Listing.cdc 0xbbbe32d615d7c84b  --network testnet```
- Send TX
```flow transactions send ./cadence/transactions/demo/sell_nft.cdc 0x0f4e1d420f06ad6d 1 10.0 --signer userA --network testnet```
- Deploy contract
```flow accounts add-contract ./cadence/contracts/UsageNFT.cdc --signer testnet-account --network testnet```
```flow accounts update-contract ./cadence/contracts/UsageNFT.cdc --signer admin --network testnet```


