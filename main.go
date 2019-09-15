package main

import (
	"fmt"
	"strconv"

	"github.com/ragul28/blockchain-go/blockchain"
)

func main() {

	bc := blockchain.InitBlock("Genesis")

	bc.AddBlock("1st block")
	bc.AddBlock("2nd block")
	bc.AddBlock("3rd block")

	for _, b := range bc.Blocks {
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)

		pow := blockchain.NewProof(b)
		fmt.Printf("pow: %s\n", strconv.FormatBool(pow.Validate()))
	}
}
