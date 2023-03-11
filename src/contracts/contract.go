package contracts

import (
	"context"
	"io/ioutil"

	constants "kxu913/demo/nft_constants"
	nft_tx "kxu913/demo/nft_tx"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/access/grpc"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/onflow/flow-go-sdk/templates"
)

func DeployContract(init bool) (string,string){
	var id string
	if(init){
		id = addContract()
	}else{
		id = updateContract()
	}
	return constants.AdminServiceAddr,id
}

func createContract() []templates.Contract {
	b, _ := ioutil.ReadFile("./cadence/contracts/AssetUsage.cdc")
	return []templates.Contract{{
		Name:   "AssetUsage",
		Source: string(b),
	}}
}

func addContract(
	
) string {
	ctx := context.Background()
	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}

	serviceAddress :=flow.HexToAddress(constants.AdminServiceAddr)
	servicePrivKey,serviceAccountKey:= nft_tx.KeyInfo(ctx,c,serviceAddress,true,"")

	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}
	tx := templates.AddAccountContract(serviceAddress, createContract()[0])
	nft_tx.SendTx(tx, serviceAddress, serviceAccountKey, 100, c, ctx, err, serviceSigner)

	// [10]
	nft_tx.WaitForSeal(ctx, c, tx.ID())
	return tx.ID().String()
}

func updateContract(
	
) string {
	ctx := context.Background()

	// [6]
	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}

	serviceAddress :=flow.HexToAddress(constants.AdminServiceAddr)
	servicePrivKey,serviceAccountKey:= nft_tx.KeyInfo(ctx,c,serviceAddress,true,"")
	serviceSigner, err := crypto.NewInMemorySigner(servicePrivKey, serviceAccountKey.HashAlgo)
	if err != nil {
		panic(err)
	}
	tx := templates.UpdateAccountContract(serviceAddress, createContract()[0])
	nft_tx.SendTx(tx, serviceAddress, serviceAccountKey, 100, c, ctx, err, serviceSigner)

	// [10]
	return tx.ID().String()
}