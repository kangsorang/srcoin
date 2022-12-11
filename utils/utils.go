package utils

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ToByte(data any) []byte {
	fmt.Println("ToByte")
	var buffer bytes.Buffer
	err := gob.NewEncoder(&buffer).Encode(data)
	HandleErr(err)
	return buffer.Bytes()
}
