import AssetUsage from 0xbbbe32d615d7c84b


transaction(asset: String, action: String) {
    prepare(user: AuthAccount) {

        if(action=="COLLECT"){
            let usage <- user.load<@AssetUsage.Usage>(from: /storage/Usage)!
            usage.collectAsset(asset: asset)
             AssetUsage.logUsage(usage: <-usage)
        }
        else if(action =="LOVE"){
            let usage <- user.load<@AssetUsage.Usage>(from: /storage/Usage)!
             usage.likeAsset(asset: asset)
             AssetUsage.logUsage(usage: <-usage)
        }


       

    }
}