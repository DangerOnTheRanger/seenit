package seenit

import (
	bolt "go.etcd.io/bbolt"
)

type BoltDatabase struct {
	boltDb *bolt.DB
}

type BoltBucket struct {
	boltDb *bolt.DB
	name string
}

func NewBoltDatabase(filename string) (*BoltDatabase, error) {
	db, err := bolt.Open(filename, 0600, nil)
	if err != nil {
		return nil, err
	}
	d := BoltDatabase{boltDb: db}
	return &d, nil
}


func (d *BoltDatabase) GetBucket(name string) (Bucket, error) {
	var bucket BoltBucket
	err := d.boltDb.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}
		bucket.boltDb = d.boltDb
		bucket.name = name
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &bucket, nil
}

func (d *BoltDatabase) Close() error {
	return d.boltDb.Close()
}

func (b *BoltBucket) Has(key string) (bool, error) {
	found := false
	err := b.boltDb.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.name))
		if bkt.Get([]byte(key)) != nil{
			found = true
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return found, nil
}

func (b *BoltBucket) Get(key string) (string, error) {
	var val string
	err := b.boltDb.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.name))
		rawVal := bkt.Get([]byte(key))
		var copyBuf []byte
		copy(copyBuf, rawVal)
		val = string(copyBuf)
		return nil
	})
	if err != nil {
		return "", err
	}
	return val, nil
}

func (b *BoltBucket) Put(key string, val string) error {
	err := b.boltDb.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(b.name))
		return bkt.Put([]byte(key), []byte(val))
	})
	return err
}
