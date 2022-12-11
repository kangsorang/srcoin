package blockchain

import (
	"crypto/sha256"
	"fmt"

	"github.com/kangsorang/srcoin/db"
	"github.com/kangsorang/srcoin/utils"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
	Height   int
}

func (b *block) persist() {
	db.SaveBlock(b.Hash, utils.ToByte(b))
}

func createBlock(data string, prevHash string, height int) *block {
	block := &block{
		Data:     data,
		PrevHash: prevHash,
		Hash:     "",
		Height:   height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}
