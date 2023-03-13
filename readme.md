## Blockchain base on Flow
*基于Flow的框架，实现一个用户在区块链上点赞功能*
### Pre condition
- *Register an account and get some flowtoken using https://testnet-faucet.onflow.org/*

- *View Account https://flow-view-source.com/testnet/account/0xbbbe32d615d7c84b/keys*

### Let start
- Deploy contract into admin account ```flow accounts add-contract ./cadence/contracts/UsageNFT.cdc --signer admin --network testnet```
- Init Account using ```flow transactions send ./cadence/transactions/init/setup_account.cdc  --signer userD --network testnet```
- Mint a NFT using ```flow transactions send ./cadence/transactions/nft/mint_nft.cdc 0x2248a8075cee5881 --signer admin --network testnet```
    - Check Mint result using ```flow scripts execute ./cadence/scripts/nft/GetIDs.cdc 0x2248a8075cee5881 --network testnet```
- Upload NFT to Storefont using ```flow transactions send ./cadence/transactions/storefont/list_nft.cdc 9 10.2 --signer userD --network testnet```
    - Check NFT in storefont using ```flow scripts execute ./cadence/scripts/storefont/GetListing.cdc 0x2248a8075cee5881  --network testnet  ```
    - Get NFT detail information in storefont using ```flow scripts execute ./cadence/scripts/storefont/GetItem.cdc 0x2248a8075cee5881 139373354  --network testnet ```
- Purchase a NFT using ```flow transactions send ./cadence/transactions/storefont/purchase_nft.cdc  0x2248a8075cee5881 139373354 --signer userB --network testnet```
    - Get NFT detail information in storefont using ```flow scripts execute ./cadence/scripts/storefont/GetItem.cdc 0x2248a8075cee5881 139373354  --network testnet ```
    - Get the NFT in userB ```flow scripts execute ./cadence/scripts/nft/GetIDs.cdc 0x0f4e1d420f06ad6d --network testnet```
*Almost done:)*

### Addtional tools
- Unlist a nft in Storefont using
    - List a nft ```flow transactions send ./cadence/transactions/storefont/list_nft.cdc 9 10.5 --signer userB --network testnet```
    - Unlist a nft ```flow transactions send ./cadence/transactions/storefont/unlist_nft.cdc 139374850  --signer userB --network testnet```
    - Clean expired listing nft ```flow transactions send ./cadence/transactions/storefont/clean_expired_items.cdc  --signer userB --network testnet```
    - Clean purchased nft ```flow transactions send ./cadence/transactions/storefont/clean_purchased_items.cdc  --signer userD --network testnet```


*Done, By the way, the demo run on flow testnet and using Flowtoken*

### Reference link
- <https://flow-view-source.com/testnet/account/0xbbbe32d615d7c84b/contract/UsageNFT>
- <https://flow-view-source.com/testnet/account/0x2d55b98eb200daef/contract/NFTStorefrontV2>
- <https://flow-view-source.com/testnet/account/0x631e88ae7f1d7c20/contract/NonFungibleToken>
- <https://flow-view-source.com/testnet/account/0x9a0766d93b6608b7/contract/FungibleToken>
- <https://flow-view-source.com/testnet/tx/4092695378644c40e4262b40624ea250776ea228476efd251d6c2e9011854670>
- <https://github.com/onflow/nft-storefront/blob/main/transactions/cleanup_purchased_listings.cdc>
- <https://developers.flow.com/cadence/tutorial/02-hello-world>




