package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/patiparnphot/simple-utxos-blockchain/blockchain"
)

type CommandLine struct {
	Blockchain *blockchain.BlockChain
}

func (cli *CommandLine) Run() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
		fmt.Println(" print - Prints the blocks in the chain")

		runtime.Goexit()
	}

	addBlockCmd := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCmd.String("block", "", "Block data")

	switch os.Args[1] {
	case "add":
		err := addBlockCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	case "print":
		err := printChainCmd.Parse(os.Args[2:])
		blockchain.Handle(err)

	default:
		fmt.Println("Usage:")
		fmt.Println(" add -block BLOCK_DATA - add a block to the chain")
		fmt.Println(" print - Prints the blocks in the chain")

		runtime.Goexit()
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			addBlockCmd.Usage()
			runtime.Goexit()
		}
		cli.Blockchain.AddBlock(*addBlockData)
		fmt.Println("Added Block!!!")
	}

	if printChainCmd.Parsed() {
		iter := cli.Blockchain.Iterator()

		for {
			block := iter.Next()

			fmt.Printf("Previous hash: %x\n", block.PrevHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Printf("Nonce: %d\n", block.Nonce)

			pow := blockchain.NewProof(block)

			fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()

			if len(block.PrevHash) == 0 {
				break
			}
		}
	}
}
