import UsageNFT from 0xbbbe32d615d7c84b
import MetadataViews from 0x631e88ae7f1d7c20
import NonFungibleToken from 0x631e88ae7f1d7c20

transaction {
    prepare(signer: AuthAccount) {
        // if the account doesn't already have a collection
        if signer.borrow<&UsageNFT.Collection>(from: UsageNFT.CollectionStoragePath) == nil {

            // create a new empty collection
            let collection <- UsageNFT.createEmptyCollection()
            
            // save it to the account
            signer.save(<-collection, to: UsageNFT.CollectionStoragePath)

            // create a public capability for the collection
            signer.link<&UsageNFT.Collection{NonFungibleToken.CollectionPublic, UsageNFT.UsageNFTCollectionPublic, MetadataViews.ResolverCollection}>(UsageNFT.CollectionPublicPath, target: UsageNFT.CollectionStoragePath)
        }else{
            signer.link<&UsageNFT.Collection{NonFungibleToken.CollectionPublic, UsageNFT.UsageNFTCollectionPublic, MetadataViews.ResolverCollection}>(UsageNFT.CollectionPublicPath, target: UsageNFT.CollectionStoragePath)
        }
    }
}