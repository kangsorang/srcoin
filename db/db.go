package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/kangsorang/srcoin/utils"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"

	checkPoint = "checkpoint"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0700, nil)
		utils.HandleErr(err)
		db = dbPointer
		err = db.Update(func(tx *bolt.Tx) error {
			_, err := tx.CreateBucketIfNotExists([]byte(dataBucket))
			utils.HandleErr(err)
			_, err = tx.CreateBucketIfNotExists([]byte(blocksBucket))
			return err
		})
		utils.HandleErr(err)

	}
	return db
}

func SaveBlock(hash string, data []byte) {
	fmt.Println("777")
	err := DB().Update(func(tx *bolt.Tx) error {
		fmt.Println("aaa")
		bucket := tx.Bucket([]byte(dataBucket))
		fmt.Println("bbb")
		err := bucket.Put([]byte(hash), data)
		fmt.Println("ccc")
		return err
	})
	utils.HandleErr(err)
	fmt.Println("888")
}

func SaveBlockchain(data []byte) {
	err := DB().Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		err := bucket.Put([]byte(checkPoint), data)
		return err
	})
	utils.HandleErr(err)
}

func GetCheckpoint() []byte {
	var data []byte
	err := DB().View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(blocksBucket))
		data = bucket.Get([]byte(checkPoint))
		return nil
	})
	utils.HandleErr(err)
	return data
}
