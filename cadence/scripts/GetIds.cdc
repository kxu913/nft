import NonFungibleToken from 0x631e88ae7f1d7c20
import UsageNFT from 0xbbbe32d615d7c84b



pub fun main(address: Address):&NonFungibleToken.NFT {
    // Get the public account object for account 0x01
    let account = getAccount(address)

    // Find the public Sale reference to their Collection
    let collectionRef = account.getCapability(UsageNFT.CollectionPublicPath)
                       .borrow<&{NonFungibleToken.CollectionPublic}>()
                       ?? panic("Could not borrow acct2 nft sale reference")
    var ids = collectionRef.borrowNFT(id:0)

    
    return ids

}
 