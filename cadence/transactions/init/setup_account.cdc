import NFTStorefrontV2 from 0x2d55b98eb200daef
import UsageNFT from 0xbbbe32d615d7c84b
import MetadataViews from 0x631e88ae7f1d7c20
import NonFungibleToken from 0x631e88ae7f1d7c20

transaction {
    prepare(acct: AuthAccount) {

        // If the account doesn't already have a Storefront
        if acct.borrow<&NFTStorefrontV2.Storefront>(from: NFTStorefrontV2.StorefrontStoragePath) == nil {

            // Create a new empty Storefront
            let storefront <- NFTStorefrontV2.createStorefront() as! @NFTStorefrontV2.Storefront
            
            // save it to the account
            acct.save(<-storefront, to: NFTStorefrontV2.StorefrontStoragePath)

            // create a public capability for the Storefront
            acct.link<&NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}>(NFTStorefrontV2.StorefrontPublicPath, target: NFTStorefrontV2.StorefrontStoragePath)
        }

        if acct.borrow<&UsageNFT.Collection>(from: UsageNFT.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- UsageNFT.createEmptyCollection()
            
            // save it to the account
            acct.save(<-collection, to: UsageNFT.CollectionStoragePath)

            // create a public capability for the collection
            acct.link<&UsageNFT.Collection{NonFungibleToken.CollectionPublic, UsageNFT.UsageNFTCollectionPublic, MetadataViews.ResolverCollection}>(UsageNFT.CollectionPublicPath, target: UsageNFT.CollectionStoragePath)
        }
    }
}