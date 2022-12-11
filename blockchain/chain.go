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
	NewestHash string
	Height     int
}

var b *blockchain
var once sync.Once

func Blockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}
			checkpoint := db.Checkpoint()
			if checkpoint == nil {
				b.addBlock("Genesis block")
				fmt.Println("addBlock end")
			} else {
				fmt.Println("Restore checkpoint")
				b.restore(checkpoint)
			}
		})
	}
	return b
}

func (b *blockchain) restore(data []byte) {
	err := gob.NewDecoder(bytes.NewReader(data)).Decode(b)
	utils.HandleErr(err)
}

func (b *blockchain) addBlock(data string) {
	fmt.Println("addBlock")
	block := createBlock(data, b.NewestHash, b.Height)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}

func (b *blockchain) persist() {
	fmt.Println("blockchain persist")
	db.SaveBlockchain(utils.ToByte(b))
}
