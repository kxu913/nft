import AssetUsage from 0xbbbe32d615d7c84b

// This transaction allows the administrator of the Voting contract
// to create a new ballot and store it in a voter's account
// The voter and the administrator have to both sign the transaction
// so it can access their storage

transaction {
    prepare(admin: AuthAccount, user: AuthAccount) {

        // borrow a reference to the admin Resource
        let adminRef = admin.borrow<&AssetUsage.Administrator>(from: /storage/Admin)!

        // create a new Ballot by calling the issueBallot
        // function of the admin Reference
        let usage <- adminRef.assignUsage()

        // store that ballot in the voter's account storage
        user.save<@AssetUsage.Usage>(<-usage, to: /storage/Usage)

        log("Ballot transferred to voter")
    }
}