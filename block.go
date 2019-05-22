package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block represents a block in the blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
	Nonce         int
}

// func (b *Block) SetHash() {
// 	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
// 	// 生成 由PrevBlockHash,data,timestamp 拼接而成的二进制数组如：[104 101 108 108 111 119 111 114 108 100]
// 	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
// 	hash := sha256.Sum256(headers)
// 	b.Hash = hash[:]
// }

// NewBlock 创建一个新的区块
func NewBlock(data string, prevBlockHash []byte) *Block { // *Block 返回一个指针变量
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0} // []byte{1,2,3} –> “[00000001 00000010 00000011]”
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// NewGenesisBlock 新建创世区块
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Serialize 序列化区块Block
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(b)
	if err != nil {
		log.Fatalln(err)
	}
	return result.Bytes()
}

// DeserializeBlock 将二制数组转换成Block
func DeserializeBlock(d []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		log.Fatalln(err)
	}
	return &block
}
