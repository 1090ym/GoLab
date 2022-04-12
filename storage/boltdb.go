package storage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
)

func SaveInstanceToDB(key []string, value []string) {
	db, err := bolt.Open("instance.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Batch(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %v", err)
		}
		for i := 0; i < len(key); i++ {
			id, _ := b.NextSequence()
			fmt.Println("id:", id)
			err = b.Put([]byte(strconv.Itoa(int(id))), []byte(value[i]))
		}
		err = b.Put([]byte("answer"), []byte("42"))
		return err
	})
}

func GetInstanceFromDB() {
	db, err := bolt.Open("instance.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Println(string(k), string(v))
		}
		return nil
	})
	fmt.Println("ending")
}

func Encode(data interface{}) bytes.Buffer {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	err := enc.Encode(data)
	if err != nil {
		return bytes.Buffer{}
	}
	return b
}

func Decode(data *bytes.Buffer, t interface{}) error {
	dec := gob.NewDecoder(data)
	err := dec.Decode(t)
	if err != nil {
		return err
	}
	return nil
}
