import FlowToken from 0x7e60df042a9c0868
import FungibleToken from 0x9a0766d93b6608b7
import NonFungibleToken from 0x631e88ae7f1d7c20
import UsageNFT from 0xbbbe32d615d7c84b
import MetadataViews from 0x631e88ae7f1d7c20
import NFTStorefrontV2 from 0x2d55b98eb200daef

pub fun getOrCreateStorefront(account: AuthAccount): &NFTStorefrontV2.Storefront {
    if let storefrontRef = account.borrow<&NFTStorefrontV2.Storefront>(from: NFTStorefrontV2.StorefrontStoragePath) {
        return storefrontRef
    }

    let storefront <- NFTStorefrontV2.createStorefront()

    let storefrontRef = &storefront as &NFTStorefrontV2.Storefront

    account.save(<-storefront, to: NFTStorefrontV2.StorefrontStoragePath)

    account.link<&NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}>(NFTStorefrontV2.StorefrontPublicPath, target: NFTStorefrontV2.StorefrontStoragePath)

    return storefrontRef
}

transaction(saleId: UFix64,saleItemPrice: UFix64) {

    // List
    let flowReceiver: Capability<&FlowToken.Vault{FungibleToken.Receiver}>
    let usageProvider: Capability<&UsageNFT.Collection{NonFungibleToken.Provider, NonFungibleToken.CollectionPublic}>
    let storefront: &NFTStorefrontV2.Storefront
    
    prepare(signer: AuthAccount) {
       
        // We need a provider capability, but one is not provided by default so we create one if needed.
        let usageNFTCollectionProviderPrivatePath = /private/usageNFTCollectionProvider

        // Receiver for the sale cut.
        self.flowReceiver = signer.getCapability<&FlowToken.Vault{FungibleToken.Receiver}>(/public/flowTokenReceiver)!
        
        assert(self.flowReceiver.borrow() != nil, message: "Missing or mis-typed FLOW receiver")

        // Check if the Provider capability exists or not if `no` then create a new link for the same.
        if !signer.getCapability<&UsageNFT.Collection{NonFungibleToken.Provider, NonFungibleToken.CollectionPublic}>(usageNFTCollectionProviderPrivatePath)!.check() {
            signer.link<&UsageNFT.Collection{NonFungibleToken.Provider, NonFungibleToken.CollectionPublic}>(usageNFTCollectionProviderPrivatePath, target: UsageNFT.CollectionStoragePath)
        }

        self.usageProvider = signer.getCapability<&UsageNFT.Collection{NonFungibleToken.Provider, NonFungibleToken.CollectionPublic}>(usageNFTCollectionProviderPrivatePath)!

        assert(self.usageProvider.borrow() != nil, message: "Missing or mis-typed KittyItems.Collection provider")

        self.storefront =  getOrCreateStorefront(account: signer)
    }

    execute {


        var totalRoyaltyCut = 0.0
        var commissionAmount = 0.0
        let effectiveSaleItemPrice = saleItemPrice-commissionAmount// commission amount is 0
        // Skip this step - Check whether the NFT implements the MetadataResolver or not.
        // Append the cut for the seller
        self.saleCuts.append(NFTStorefrontV2.SaleCut(
            receiver: self.flowReceiver,
            amount: effectiveSaleItemPrice - totalRoyaltyCut
        ))
        
        // Execute to create listing
        self.storefront.createListing(
            nftProviderCapability: self.usageProvider,
            nftType: Type<@UsageNFT.NFT>(),
            nftID: saleId,
            salePaymentVaultType: Type<@FlowToken.Vault>(),
            saleCuts: self.saleCuts,
            marketplacesCapability: self.marketplacesCapability.length == 0 ? nil : self.marketplacesCapability,
            customID: nil,
            commissionAmount: UFix64(0),
            expiry: UInt64(getCurrentBlock().timestamp) + UInt64(500)
        )
    }

    post {
        self.recipientCollectionRef.getIDs().contains(self.mintingIDBefore): "The next NFT ID should have been minted and delivered"
        UsageNFT.totalSupply == self.mintingIDBefore + 1: "The total supply should have been increased by 1"
    }
}
 