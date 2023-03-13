/*
*
*   In this example, we want to create a simple approval voting contract
*   where a polling place issues ballots to addresses.
*
*   The run a vote, the Admin deploys the smart contract,
*   then initializes the proposals
*   using the initialize_proposals.cdc transaction.
*   The array of proposals cannot be modified after it has been initialized.
*
*   Then they will give ballots to users by
*   using the issue_ballot.cdc transaction.
*
*   Every user with a ballot is allowed to approve any number of proposals.
*   A user can choose their votes and cast them
*   with the cast_vote.cdc transaction.
*
*/

pub contract AssetUsage {

    //list of proposals to be approved
    pub var assets: [String]

    // number of votes per proposal
    pub let likes: {String: Int}

    pub let collects: {String: Int}

    // This is the resource that is issued to users.
    // When a user gets a Ballot object, they call the `vote` function
    // to include their votes, and then cast it in the smart contract
    // using the `cast` function to have their vote included in the polling
    pub resource Usage {

        // array of all the proposals
        pub let assets: [String]

        // corresponds to an array index in proposals after a vote
        pub var likes: {String:Int}

        pub var collects: {String:Int}

        init() {
            self.assets = AssetUsage.assets
            self.likes = {}
            self.collects = {}
            var i = 0
            
            while i < self.assets.length {
                self.likes[self.assets[i]] = 0
                self.collects[self.assets[i]] = 0
                i = i + 1
            }
        }

        // modifies the ballot
        // to indicate which proposals it is voting for
        pub fun likeAsset(asset: String) {
            pre {
                AssetUsage.assets.contains(asset): "Cannot like an unexist asset for a proposal that doesn't exist"
            }
            self.likes[asset]=self.likes[asset]!+1
        }

        pub fun collectAsset(asset: String) {
            pre {
                AssetUsage.assets.contains(asset): "Cannot like an unexist asset for a proposal that doesn't exist"
            }
            self.collects[asset] = self.collects[asset]!+1
        }
    }

    // Resource that the Administrator of the vote controls to
    // initialize the proposals and to pass out ballot resources to voters
    pub resource Administrator {

        // function to initialize all the proposals for the voting
        pub fun addAsset(_ asset: String) {
            pre {
                asset !="": "Cannot add an empty asset"
            }
            AssetUsage.assets.append(asset)
            AssetUsage.likes.insert(key: asset, 0)
            AssetUsage.collects.insert(key: asset, 0)

           
        }

        // The admin calls this function to create a new Ballo
        // that can be transferred to another user
        pub fun assignUsage(): @Usage {
            return <-create Usage()
        }
    }

    // A user moves their ballot to this function in the contract where
    // its votes are tallied and the ballot is destroyed
    pub fun logUsage(usage: @Usage) {
        var i =0;
        while i< self.assets.length {
        var id = self.assets[i];
            if usage.likes[id]==1{
                self.likes[id] =  self.likes[id]!+1
            }
            if usage.collects[id]==1{
                self.collects[id] =  self.collects[id]!+1
            }
            i = i+1
        }
        destroy usage
    }

    // initializes the contract by setting the proposals and votes to empty
    // and creating a new Admin resource to put in storage
    init() {
        self.assets = []
        self.likes = {}
        self.collects ={}

        self.account.save<@Administrator>(<-create Administrator(), to: /storage/Admin)
    }
}