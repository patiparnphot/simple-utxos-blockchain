package main

import (
	"os"

	"github.com/patiparnphot/simple-utxos-blockchain/blockchain"
	"github.com/patiparnphot/simple-utxos-blockchain/cli"
)

func main() {
	defer os.Exit(0)

	chain := blockchain.InitBlockChain()
	defer chain.Database.Close()

	cmd := cli.CommandLine{Blockchain: chain}
	cmd.Run()

}
