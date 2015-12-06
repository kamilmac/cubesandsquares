package main

import (
    "fmt"
    "log"
    "time"
    "github.com/satori/go.uuid"
	"github.com/boltdb/bolt"
)

func getUid() (id string) {
    return uuid.NewV4().String()
}

func openDb(path string) (DB *bolt.DB) {
	DB, err := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	return
}

func put(bucket string, id string, value []byte) {
    err := db.Update(func(tx *bolt.Tx) error {
        b, err := tx.CreateBucketIfNotExists([]byte(bucket))
        if err != nil {
            return fmt.Errorf("create bucket: %s", err)
        }
        b.Put([]byte(id), []byte(value))
        return nil
    })
    if err != nil {
		log.Fatal(err)
	}
}

func get(bucket, key string) (v []byte) {
    db.View(func(tx *bolt.Tx) error {
	    b := tx.Bucket([]byte(bucket))
        v = b.Get([]byte(key))
	    return nil
	})
	return
}

func delete(bucket, key string) {
    db.Update(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(bucket))
        b.Delete([]byte(key))
        return nil
    })
}

func getAll(bucket string) (list []string){
    db.View(func(tx *bolt.Tx) error {

	    b := tx.Bucket([]byte(bucket))
        
        if(b != nil) {
            b.ForEach(func(k, v []byte) error {
                list = append(list, string(v))
                return nil
            })
        }
        
	    return nil
	})
	return
}