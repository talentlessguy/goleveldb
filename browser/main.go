//go:build js
// +build js

package main

import (
	"fmt"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, err := leveldb.OpenFile("/tmp/test", nil)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}
	defer db.Close()

	db.Put([]byte("key"), []byte("value"), nil)
	fmt.Println("Put key:value")
	val, err := db.Get([]byte("key"), nil)
	if err != nil {
		fmt.Println("Error getting value:", err)
		os.Exit(1)
	}
	fmt.Println("Get key:value", string(val))

	os.Exit(0)
}
