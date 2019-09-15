package blockchain

type Blockchain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// Create & Add Blocks to blockchain
func CreateBlock(Data string, PrevHash []byte) *Block {

	block := &Block{[]byte{}, []byte(Data), PrevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

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
