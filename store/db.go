package store

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"

	"go.etcd.io/bbolt"
)

var (
	bucketName = []byte("items")
)

type DB struct {
	conn *bbolt.DB
}

// Open opens the database.
func Open(dbFile string) (*DB, error) {
	conn, err := bbolt.Open(dbFile, 0666, nil)
	if err != nil {
		return nil, err
	}

	// Make sure bucket exists
	err = conn.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucket(bucketName)
		if !errors.Is(err, bbolt.ErrBucketExists) {
			return err
		}
		return nil
	})

	if err != nil {
		conn.Close()
		return nil, err
	}

	db := DB{
		conn: conn,
	}
	return &db, nil
}

func (db *DB) Close() error {
	return db.conn.Close()
}

func (db *DB) Update(i Item) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(i); err != nil {
		return err
	}

	err := db.conn.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			bname := string(bucketName)
			return fmt.Errorf("%q: bucket not found", bname)
		}

		key := []byte(i.SKU)
		return b.Put(key, buf.Bytes())
	})

	return err
}

var ErrNotFound = errors.New("not found")

func (db *DB) Get(sku string) (Item, error) {
	var data []byte

	err := db.conn.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(bucketName)
		if b == nil {
			bname := string(bucketName)
			return fmt.Errorf("%q: bucket not found", bname)
		}

		data = b.Get([]byte(sku))
		return nil
	})

	if err != nil {
		return Item{}, err
	}

	if data == nil {
		return Item{}, ErrNotFound
	}

	var i Item
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(&i); err != nil {
		return Item{}, nil
	}

	return i, nil
}
