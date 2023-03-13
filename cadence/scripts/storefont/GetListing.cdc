import NFTStorefrontV2 from 0x2d55b98eb200daef
pub fun main(account: Address): [UInt64]{
    let storefrontRef = getAccount(account)
        .getCapability<&NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}>(
            NFTStorefrontV2.StorefrontPublicPath
        )
        .borrow()
        ?? panic("Could not borrow public storefront from address")

    return storefrontRef.getListingIDs()
}