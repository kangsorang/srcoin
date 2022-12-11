package db

import (
	"github.com/boltdb/bolt"
	"github.com/kangsorang/srcoin/utils"
)

const (
	database    = "blockchain.db"
	blockBucket = "block"
	dataBucket  = "data"
	checkpoint  = "checkpoint"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPoint, err := bolt.Open(database, 0700, nil)
		utils.HandleErr(err)
		db = dbPoint
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(blockBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(dataBucket))
			return err
		})
		utils.HandleErr(err)
	}
	return db
}

func SaveBlock(hash string, data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blockBucket))
		err := bucket.Put([]byte(hash), data)
		return err
	})
	utils.HandleErr(err)
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		err := bucket.Put([]byte(checkpoint), data)
		return err
	})
	utils.HandleErr(err)
}

func Checkpoint() []byte {
	var checkpoint []byte
	DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(dataBucket))
		checkpoint = bucket.Get([]byte(checkpoint))
		return nil
	})
	return checkpoint
}
