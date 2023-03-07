package main

import (
	"flag"
	"fmt"
	abciClient "github.com/tendermint/tendermint/abci/client"
	"github.com/tendermint/tendermint/abci/types"
)

var clientAddr string

func init() {
	flag.StringVar(&clientAddr, "client-addr", "tcp://0.0.0.0:26658", "Unix domain client address")
}

func main() {
	socketClient := abciClient.NewSocketClient(clientAddr, false)

	if err := socketClient.Start(); err != nil {
		fmt.Println(err.Error())
	}

	res, err := socketClient.ListSnapshotsSync(types.RequestListSnapshots{})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(res.Snapshots)

	resChunk, errChunk := socketClient.LoadSnapshotChunkSync(types.RequestLoadSnapshotChunk{
		Height: 6000,
		Format: 2,
		Chunk:  0,
	})
	if errChunk != nil {
		fmt.Println(errChunk.Error())
	}

	fmt.Println(resChunk)
}
