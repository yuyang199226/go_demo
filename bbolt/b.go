package main

import (
    "log"

    bolt "go.etcd.io/bbolt"
)

func main() {
    // Open the my.db data file in your current directory.
    // It will be created if it doesn't exist.
    db, err := bolt.Open("my.db", 0600, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    tx,err := db.Begin(true)
    if err != nil {
        log.Fatal(err)
    }
    defer tx.Rollback()
    _, err = tx.CreateBucket([]byte("myBucket"))
    if err != nil {
        log.Fatal(err)
    }
    if err := tx.Commit();err != nil {
        log.Fatal(err)
    }
}
