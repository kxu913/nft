import NonFungibleToken from 0x631e88ae7f1d7c20
import UsageNFT from 0xbbbe32d615d7c84b


transaction(recipient: Address,saleItemPrice: UFix64) {

    let minter: &UsageNFT.NFTMinter

    let recipientCollectionRef: &{NonFungibleToken.CollectionPublic}

    
    let mintingIDBefore: UInt64

    prepare(signer: AuthAccount) {

        self.mintingIDBefore = UsageNFT.totalSupply

        self.minter = signer.borrow<&UsageNFT.NFTMinter>(from: UsageNFT.MinterStoragePath)
            ?? panic("Could not borrow a reference to the NFT minter")
        self.recipientCollectionRef = getAccount(recipient)
            .getCapability(UsageNFT.CollectionPublicPath)
            .borrow<&{NonFungibleToken.CollectionPublic}>()
            ?? panic("Could not get receiver reference to the NFT Collection")
    }

    execute {
        self.minter.mintNFT(
            recipient: self.recipientCollectionRef,
            name: "demo",
            description: "demo desc",
            thumbnail: "",
            royalties: []
        )

    }
    post {
        self.recipientCollectionRef.getIDs().contains(self.mintingIDBefore): "The next NFT ID should have been minted and delivered"
        UsageNFT.totalSupply == self.mintingIDBefore + 1: "The total supply should have been increased by 1"
    }

}
 