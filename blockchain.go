package main

import (
	"log"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"

// BlockChain 区块链
type BlockChain struct {
	tip []byte
	db  *bolt.DB
}

// AddBlock 添加新的区块
func (bc *BlockChain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newblock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newblock)
}

// NewBlockchain creates a new Blockchain with genesis Block
// Open a DB file.
// Check if there’s a blockchain stored in it.
// If there’s a blockchain:
//     Create a new Blockchain instance.
// 		 Set the tip of the Blockchain instance to the last block hash stored in the DB.
// If there’s no existing blockchain:
// 		 Create the genesis block.
// 		 Store in the DB.
// 		 Save the genesis block’s hash as the last block hash.
// 		 Create a new Blockchain instance with its tip pointing at the genesis block.
func NewBlockchain() *BlockChain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket)) // []byte(blocksBucket) 生成blocksBucket 的byte数组
			err = b.Put(genesis.Hash, genesis.Serialize())
			err = b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("1"))
		}
		return nil
	})
	bc := BlockChain{tip, db}
	return &bc
}
