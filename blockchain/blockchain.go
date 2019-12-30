package blockchain

import (
	"fmt"
	"github.com/dgraph-io/badger"
	"log"
)

const (
	dbpath = "./db/blocks"
)

type Blockchain struct {
	LastHash []byte
	DB       *badger.DB
}

type BlockchainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// Initialize Blockchain with Genesis
func InitBlock(Data string) *Blockchain {
	var lastHash []byte

	db, err := badger.Open(badger.DefaultOptions(dbpath))
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	// db start with rw transactions
	err = db.Update(func(txn *badger.Txn) error {
		// Check db for exisiting block
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("Genesis Block")

			// Create Genesis block
			genesis := CreateBlock(Data, []byte{})
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)

			// Record lasthash in db
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err

		} else {
			// Create chain block
			item, err := txn.Get([]byte("lh"))
			Handle(err)
			err = item.Value(func(lhValue []byte) error {
				lastHash = append([]byte{}, lhValue...)
				return nil
			})
			return err
		}
	})

	return &Blockchain{lastHash, db}
}

// Create & Add Blocks to blockchain
func (c *Blockchain) AddBlock(data string) {
	var lastHash []byte

	err := c.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(lhValue []byte) error {
			lastHash = append([]byte{}, lhValue...)
			return err
		})
		return err
	})
	Handle(err)

	newblock := CreateBlock(data, lastHash)

	err = c.DB.Update(func(txn *badger.Txn) error {
		err := txn.Set(newblock.Hash, newblock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newblock.Hash)

		c.LastHash = newblock.Hash

		return err
	})
	Handle(err)
}

func (c *Blockchain) Iterator() *BlockchainIterator {
	itr := &BlockchainIterator{c.LastHash, c.DB}
	return itr
}

func (itr *BlockchainIterator) Next() *Block {
	var block *Block
	var encodedBk []byte

	err := itr.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(itr.CurrentHash)
		Handle(err)
		err = item.Value(func(lhValue []byte) error {
			encodedBk = append([]byte{}, lhValue...)
			return err
		})
		block = Deserialize(encodedBk)
		return err
	})
	Handle(err)

	itr.CurrentHash = block.PrevHash
	return block
}
