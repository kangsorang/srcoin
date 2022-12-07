package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	Blocks []*Block
}

var bc *blockchain
var once sync.Once

// get blockchain
func GetBlockchain() *blockchain {
	if bc == nil {
		once.Do(func() {
			bc = &blockchain{}
			bc.AddBlock("Genesis block")
		})
	}
	return bc
}

// create new block
func createBlock(data string) *Block {
	newBlock := Block{Data: data, Hash: "", PrevHash: getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// AddBlock to blockchain
func (bc *blockchain) AddBlock(data string) {
	bc.Blocks = append(bc.Blocks, createBlock(data))
}

func (bc *blockchain) AllBlock() []*Block {
	return bc.Blocks
}

// calculate block hash
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(getLastHash() + b.Data))
	b.Hash = fmt.Sprintf("%x", hash)
}

// get last block hash
func getLastHash() string {
	blockLength := len(GetBlockchain().Blocks)
	if blockLength == 0 {
		return ""
	}
	return GetBlockchain().Blocks[blockLength-1].Hash
}
