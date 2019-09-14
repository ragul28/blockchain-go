package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type Blockchain struct {
	Blocks []*Block
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

// Create & Add Blocks to blockchain
func CreateBlock(Data string, PrevHash []byte) *Block {

	hash := HashGen(Data, PrevHash)
	block := &Block{hash, []byte(Data), PrevHash}

	return block
}

// Initialize Blockchain with Genesis
func InitBlock(Data string) *Blockchain {

	Genesis := CreateBlock(Data, []byte{})
	return &Blockchain{[]*Block{Genesis}}
}

// Create & Add Blocks to blockchain
func (c *Blockchain) AddBlock(Data string) {

	ph := c.Blocks[len(c.Blocks)-1].Hash
	block := CreateBlock(Data, ph)
	c.Blocks = append(c.Blocks, block)
}
