package transactions

import (
	"context"
	"fmt"
	"io/ioutil"
	. "kxu913/demo/nft_constants"
	"time"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
)

func ExecuteTransacation(transicationName string, recipientAddress string,asset string) string {
	id := newTransacation(getTransacationFile(transicationName), recipientAddress, asset, 100)
	return id
}

func ExecuteTransacationUsingMultipleSig(transicationName string, recipientAddress string,privateKey string,asset string) string {
	id := newTransacationUsingMultipleUser(getTransacationFile(transicationName), recipientAddress, privateKey, asset, 100)
	return id
}

func ExecuteUserTransacation(transicationName string, address string,privateKey string, asset string) string {
	id := newUserTransacation(getTransacationFile(transicationName), address, privateKey,asset, 100)
	return id
}

func KeyInfo(ctx context.Context,c *grpc.Client, serviceAddress flow.Address,isAdmin bool,privateKey string)(crypto.PrivateKey,*flow.AccountKey){
	serviceSigAlgo := crypto.StringToSignatureAlgorithm(SigAlgoName)

	if(isAdmin){
		privateKey = AdminServicePrivKey
	}
	servicePrivKey, err := crypto.DecodePrivateKeyHex(serviceSigAlgo, privateKey)
	if err != nil {
		panic(err)
	}
	serviceAccount, err := c.GetAccountAtLatestBlock(ctx, serviceAddress)
	if err != nil {
		panic(err)
	}
	
	return servicePrivKey,serviceAccount.Keys[0]
}

func KeyInfo2(ctx context.Context,c *grpc.Client, serviceAddress flow.Address, privateKey string)(crypto.PrivateKey,*flow.AccountKey){
	sigAlgo := crypto.StringToSignatureAlgorithm(SigAlgoName)

	priKey, err := crypto.DecodePrivateKeyHex(sigAlgo, privateKey)
	if err != nil {
		panic(err)
	}
	account, err := c.GetAccountAtLatestBlock(ctx, serviceAddress)

	if err != nil {
		panic(err)
	}
	
	return priKey,account.Keys[0]
}

func getTransacationFile(transicationName string) string {
	return "./cadence/transactions/" + transicationName + ".cdc"
}

func SendTx(tx *flow.Transaction, serviceAddress flow.Address, serviceAccountKey *flow.AccountKey, gasLimit uint64, c *grpc.Client, ctx context.Context, err error, serviceSigner crypto.InMemorySigner) {
	tx.SetProposalKey(serviceAddress, 0, serviceAccountKey.SequenceNumber)
	tx.SetPayer(serviceAddress)
	tx.SetGasLimit(uint64(gasLimit))
	latestBlock, _ := c.GetLatestBlock(ctx, true)
	tx.SetReferenceBlockID(latestBlock.ID)

	err = tx.SignEnvelope(serviceAddress, 0, serviceSigner)
	if err != nil {
		panic(err)
	}

	err = c.SendTransaction(ctx, *tx)
	if err != nil {
		panic(err)
	}
}

func WaitForSeal(ctx context.Context, c *grpc.Client, id flow.Identifier) *flow.TransactionResult {
	result, err := c.GetTransactionResult(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Waiting for transaction %s to be sealed...\n", id)

	for result.Status != flow.TransactionStatusSealed {
		time.Sleep(time.Second)
		fmt.Print(".")
		result, err = c.GetTransactionResult(ctx, id)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println()
	fmt.Printf("Transaction %s sealed\n", id)
	return result
}

func newTransacation(
	fileName string,
	recipientAddressHex string,
	asset string,
	gasLimit uint64,
) string {
	ctx := context.Background()
	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}
	serviceAddress :=flow.HexToAddress(AdminServiceAddr)
	servicePrivKey,serviceAccountKey:= KeyInfo(ctx,c,serviceAddress,true,"")

	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}

	s, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tx := flow.NewTransaction().
		SetScript(s).AddAuthorizer(flow.HexToAddress(AdminServiceAddr))
	if recipientAddressHex != "" {
		tx.AddAuthorizer(flow.HexToAddress(recipientAddressHex))
	}
	if asset !="" {
		a,err := cadence.NewString(asset)
		if err!=nil {
			panic(err)
		}
		tx.AddArgument(a)
	}

	SendTx(tx, serviceAddress, serviceAccountKey, gasLimit, c, ctx, err, serviceSigner)
	WaitForSeal(ctx, c, tx.ID())

	// [10]
	return tx.ID().String()
}


func newTransacationUsingMultipleUser(
	fileName string,
	recipientAddressHex string,
	privateKey string,
	asset string,
	gasLimit uint64,
) string {
	ctx := context.Background()
	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}

	serviceAddress :=flow.HexToAddress(AdminServiceAddr)
	servicePrivKey,serviceAccountKey:= KeyInfo2(ctx,c,serviceAddress,AdminServicePrivKey)
	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}

	acct2Adress :=flow.HexToAddress(recipientAddressHex)
	acct2PrivKey,acct2AccountKey:= KeyInfo2(ctx,c,acct2Adress,privateKey)
	acct2Signer, err := crypto.NewInMemorySigner(acct2PrivKey, acct2AccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}

	s, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	
	latestBlock, _ := c.GetLatestBlock(ctx, true)
	// tx := flow.NewTransaction().
	// SetScript(s).
	// AddAuthorizer(serviceAddress)

	// tx.SetProposalKey(serviceAddress, 0, serviceAccountKey.SequenceNumber)
	// tx.SetPayer(serviceAddress)
	// tx.SetGasLimit(uint64(gasLimit))
	// // latestBlock, _ := c.GetLatestBlock(ctx, true)
	// tx.SetReferenceBlockID(latestBlock.ID)

	// err = tx.SignEnvelope(serviceAddress, 0, serviceSigner)
	// if err != nil {
	// 	panic(err)
	// }

	// err = c.SendTransaction(ctx, *tx)
	// if err != nil {
	// 	panic(err)
	// }


	tx := flow.NewTransaction().
			SetScript(s).
			AddAuthorizer(serviceAddress).
			AddAuthorizer(acct2Adress)
	
	tx.SetProposalKey(serviceAddress, 0, serviceAccountKey.SequenceNumber)
	tx.SetPayer(acct2Adress)
	tx.SetGasLimit(uint64(gasLimit))
	tx.SetReferenceBlockID(latestBlock.ID)
	//https://github.com/onflow/flow/blob/master/docs/content/concepts/accounts-and-keys.md#payload
	//be sure that payer address must be sign off at last
	
	err = tx.SignPayload(serviceAddress,0, serviceSigner)
	if err!=nil{
		panic(err)
	}
	err = tx.SignEnvelope(acct2Adress,0, acct2Signer)
	if err!=nil{
		panic(err)
	}

	// if asset !="" {
	// 	a,err := cadence.NewString(asset)
	// 	if err!=nil {
	// 		panic(err)
	// 	}
	// 	tx.AddArgument(a)
	// }

	err = c.SendTransaction(ctx, *tx)
	if err != nil {
		panic(err)
	}

	WaitForSeal(ctx, c, tx.ID())
	return tx.ID().String()
}

func newUserTransacation(
	fileName string,
	recipientAddressHex string,
	privateKey string,
	asset string,
	gasLimit uint64,
) string {
	ctx := context.Background()
	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}
	serviceAddress :=flow.HexToAddress(recipientAddressHex)
	servicePrivKey,serviceAccountKey:= KeyInfo(ctx,c,serviceAddress,false,privateKey)

	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}

	s, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	tx := flow.NewTransaction().
		SetScript(s).AddAuthorizer(flow.HexToAddress(recipientAddressHex))

	if asset !="" {
		a,err := cadence.NewString(asset)
		if err!=nil {
			panic(err)
		}
		tx.AddArgument(a)

		l,err := cadence.NewString("LOVE")
		if err!=nil {
			panic(err)
		}
		tx.AddArgument(l)
	}

	SendTx(tx, serviceAddress, serviceAccountKey, gasLimit, c, ctx, err, serviceSigner)
	WaitForSeal(ctx, c, tx.ID())

	// [10]
	return tx.ID().String()
}

