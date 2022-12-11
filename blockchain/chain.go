package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/kangsorang/srcoin/db"
	"github.com/kangsorang/srcoin/utils"
)

type blockchain struct {
	NewestHash string `json:"NewestHash"`
	Height     int    `json:"height"`
}

var b *blockchain
var once sync.Once

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkpoint := db.GetCheckpoint()
			fmt.Println("111")
			if checkpoint == nil {
				b.AddBlock("Genesis")
			} else {
				fmt.Println("restore checkpoint")
				b.restore(checkpoint)
			}
			fmt.Println("222")
		})
	}
	return b
}

func (b *blockchain) restore(data []byte) {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(b)
	utils.HandleErr(err)
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height+1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	fmt.Println("222")
	b.persist()
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}
