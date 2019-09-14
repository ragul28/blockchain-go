package main

import (
	"fmt"

	"github.com/ragul28/blockchain-go/blockchain"
)

func main() {

	blockchain := blockchain.InitBlock("Genesis")

	blockchain.AddBlock("1st block")
	blockchain.AddBlock("2nd block")
	blockchain.AddBlock("3rd block")

	for _, b := range blockchain.Blocks {
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
	}
}
