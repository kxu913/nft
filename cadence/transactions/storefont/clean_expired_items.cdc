
import NFTStorefrontV2 from 0x2d55b98eb200daef

transaction() {
    let storefront: &NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}

    prepare(acct: AuthAccount) {
        self.storefront = acct
            .getCapability<&NFTStorefrontV2.Storefront{NFTStorefrontV2.StorefrontPublic}>(
                NFTStorefrontV2.StorefrontPublicPath
            )!
            .borrow()
            ?? panic("Could not borrow Storefront from provided address")
    }

    execute {
        // Be kind and recycle
        var ids = self.storefront.getListingIDs()
        if(ids.length>0){
            self.storefront.cleanupExpiredListings(fromIndex: 0, toIndex:UInt64(ids.length))
        }
        
    }
}