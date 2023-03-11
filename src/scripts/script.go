package scripts

import (
	"context"

	"io/ioutil"

	"github.com/onflow/cadence"
	"github.com/onflow/flow-go-sdk/access/grpc"
)

func GetUsage() string{

	ctx := context.Background()

	c, err := grpc.NewClient(grpc.TestnetHost)
	if err != nil {
		panic("failed to connect to node")
	}
	s, err := ioutil.ReadFile(getScriptFile())
	if err != nil {
		panic(err)
	}

	// argAddress :=cadence.NewAddress(flow.HexToAddress(AdminServiceAddr))
	args := []cadence.Value{}
	// args = append(args, argAddress)

	value, err := c.ExecuteScriptAtLatestBlock(ctx, s, args)
	if err != nil {
		panic(err)
	}

	return value.String()
}

func getScriptFile() string {
	return "./cadence/scripts/GetUsage.cdc"
}