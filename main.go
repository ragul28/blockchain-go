package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Blockchain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// Generate Hash for block
func HashGen(Data string, PrevHash []byte) []byte {
	info := bytes.Join([][]byte{[]byte(Data), PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	return hash[:]
}

// Initialize Blockchain with Genesis
func InitBlock(Data string) *Blockchain {
	Genesishash := HashGen(Data, []byte{})
	Genesis := &Block{Genesishash, []byte("Genesis"), []byte{}}

	return &Blockchain{[]*Block{Genesis}}
}

// Create & Add Blocks to blockchain
func (c *Blockchain) AddBlock(Data string) {

	ph := c.blocks[len(c.blocks)-1].Hash
	hash := HashGen(Data, ph)

	block := &Block{hash, []byte(Data), ph}

	c.blocks = append(c.blocks, block)
}

func main() {

	blockchain := InitBlock("Genesis")

	blockchain.AddBlock("1st block")
	blockchain.AddBlock("2nd block")
	blockchain.AddBlock("3rd block")

	for _, b := range blockchain.blocks {
		fmt.Printf("Data: %s\n", b.Data)
		fmt.Printf("Hash: %x\n", b.Hash)
		fmt.Printf("PrevHash: %x\n", b.PrevHash)
	}
}
