package main

import (
    "fmt"
    "log"
    "time"
    "github.com/satori/go.uuid"
	"github.com/boltdb/bolt"
)

func getUid() (id []byte) {
    return uuid.NewV4().Bytes()
}

func openDb(path string) (DB *bolt.DB) {
	DB, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func put(bucket, id, value []byte) {
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists(bucket)
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        b.Put([]byte(id), value)
        return nil
    })
    if err != nil {
		log.Fatal(err)
	}
}

func get(bucket, key []byte) (v []byte) {
    db.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket(bucket)
        v = b.Get(key)
	    return nil
	})
	return
}

func getAll(bucket []byte) (list [][]byte){
	db.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte(bucket))
	    b.ForEach(func(k, v []byte) error {
            list = append(list, v)
	        return nil
	    })
	    return nil
	})
	return
}