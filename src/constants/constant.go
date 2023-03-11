package constants

import (
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
)
var (
	SigAlgoName  = crypto.ECDSA_P256.String()
	HashAlgoName = crypto.SHA3_256.String()
	AdminServiceAddr = "0xbbbe32d615d7c84b"
	AdminServicePrivKey = "privateKey"
	AdminServicePubKey = "publicKey"
	Host              = grpc.TestnetHost

)