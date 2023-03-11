import AssetUsage from 0xbbbe32d615d7c84b

transaction(asset: String){

 prepare(admin: AuthAccount) {
  let adminRef = admin.borrow<&AssetUsage.Administrator>(from: /storage/Admin)!
  adminRef.addAsset(asset)
 }
}