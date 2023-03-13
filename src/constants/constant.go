package constants

import (
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
)
var (
	SigAlgoName  = crypto.ECDSA_P256.String()
	HashAlgoName = crypto.SHA3_256.String()
	AdminServiceAddr = "0xbbbe32d615d7c84b"
	AdminServicePrivKey = "21eda1886371b57badd1812fb7dddb489a37c95625c015eacd6dd59d3fe0172b"
	AdminServicePubKey = "cc702f779ba96b015e25f3c98ee537b540c55d2417e7be5ff4564c0af09f26f485c1a26bf8da0c3ca07741cc08fb45a4e0adb2c6c8929f8276df0dd632226a77"
	Host              = grpc.TestnetHost

)