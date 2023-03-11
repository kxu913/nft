package account

import (
	"context"
	"crypto/rand"
	"encoding/hex"

	constants "kxu913/demo/nft_constants"

	nft_tx "kxu913/demo/nft_tx"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
	templates "github.com/onflow/flow-go-sdk/templates"
)

type Account struct{
	Address string
	PublicKey string
	PrivateKey string
}

func CreateAccount() *Account{
	publicKeyHex, privateKeyHex := generateKeys(constants.SigAlgoName)

	ctx := context.Background()

	sigAlgo := crypto.StringToSignatureAlgorithm(constants.SigAlgoName)
	publicKey, err := crypto.DecodePublicKeyHex(sigAlgo, publicKeyHex)
	if err != nil {
		panic(err)
	}

	hashAlgo := crypto.StringToHashAlgorithm(constants.HashAlgoName)

	// [4]
	accountKey := flow.NewAccountKey().
		SetPublicKey(publicKey).
		SetSigAlgo(sigAlgo).
		SetHashAlgo(hashAlgo).
		SetWeight(flow.AccountKeyWeightThreshold)

	// // [5]
	// accountCode := []byte(nil)
	// if strings.TrimSpace(code) != "" {
	// 	accountCode = []byte(code)
	// }

	// [6]
	c, err := grpc.NewClient(constants.Host)
	if err != nil {
		panic("failed to connect to node")
	}

	serviceSigAlgo := crypto.StringToSignatureAlgorithm(constants.SigAlgoName)
	servicePrivKey, err := crypto.DecodePrivateKeyHex(serviceSigAlgo,constants. AdminServicePrivKey)
	if err != nil {
		panic(err)
	}

	serviceAddress := flow.HexToAddress(constants.AdminServiceAddr)
	serviceAccount, err := c.GetAccountAtLatestBlock(ctx, serviceAddress)
	if err != nil {
		panic(err)
	}

	// [7]
	serviceAccountKey := serviceAccount.Keys[0]
	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}
	// [8]
	tx, err := templates.CreateAccount([]*flow.AccountKey{accountKey}, nil, serviceAddress)
	if err != nil {
		panic(err)
	}
	nft_tx.SendTx(tx, serviceAddress, serviceAccountKey, 100, c, ctx, err, serviceSigner)
	// [10]

	res := nft_tx.WaitForSeal(ctx, c, tx.ID())
	address := ""

	for _, e := range res.Events {
		if e.Type == flow.EventAccountCreated {
			accountCreatedEvent := flow.AccountCreatedEvent(e)
			address = accountCreatedEvent.Address().Hex()
		}
	}

	return &Account{
		Address: address,
		PrivateKey: privateKeyHex,
		PublicKey: publicKeyHex,
	}
}


func generateKeys(sigAlgoName string) (string, string) {
	seed := make([]byte, crypto.MinSeedLength)
	_, err := rand.Read(seed)
	if err != nil {
		panic(err)
	}

	// [3]
	sigAlgo := crypto.StringToSignatureAlgorithm(sigAlgoName)
	privateKey, err := crypto.GeneratePrivateKey(sigAlgo, seed)
	if err != nil {
		panic(err)
	}

	// [4]
	publicKey := privateKey.PublicKey()

	pubKeyHex := hex.EncodeToString(publicKey.Encode())
	privKeyHex := hex.EncodeToString(privateKey.Encode())

	return pubKeyHex, privKeyHex
}