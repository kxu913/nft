
import NFTStorefrontV2 from 0x2d55b98eb200daef

transaction(fromIndex: UInt64, toIndex: UInt64, storefrontAddress: Address) {
    let storefront: &NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}

    prepare(acct: AuthAccount) {
        self.storefront = getAccount(storefrontAddress)
            .getCapability<&NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}>(
                NFTStorefrontV2.StorefrontPublicPath
            )!
            .borrow()
            ?? panic("Could not borrow Storefront from provided address")
    }

    execute {
        // Be kind and recycle
        self.storefront.cleanupExpiredListings(fromIndex: fromIndex, toIndex: toIndex)
    }
}