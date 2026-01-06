package main

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"go.etcd.io/bbolt"
)

var (
	bucketName = []byte("rides")
)

type DB struct {
	conn *bbolt.DB
}

func NewDB(fileName string) (*DB, error) {
	conn, err := bbolt.Open(fileName, 0666, nil)
	if err != nil {
		return nil, err
	}

	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	})

	if err != nil {
		conn.Close()
		return nil, err
	}

	db := DB{conn}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Insert(r Ride) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(r); err != nil {
		return err
	}

	err := db.conn.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		return b.Put([]byte(r.ID), buf.Bytes())
	})
	return err
}

func (db *DB) Get(id string) (Ride, error) {
	var r Ride

	err := db.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		data := b.Get([]byte(id))
		if data == nil {
			return fmt.Errorf("%q not found", id)
		}

		if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&r); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return Ride{}, err
	}

	return r, nil
}
