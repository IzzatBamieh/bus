package db

import (
	"os"

	"github.com/dgraph-io/badger"
)

type OffsetStore struct {
	id string
	db *badger.DB
}

func NewOffsetStore(id string, workdir string) (*OffsetStore, error) {
	if err := os.MkdirAll(workdir, 0700); err != nil {
		return nil, err
	}
	opts := badger.DefaultOptions
	opts.Dir = workdir
	opts.ValueDir = workdir
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return &OffsetStore{
		id,
		db,
	}, nil
}

func (store *OffsetStore) Get(key []byte) ([]byte, error) {
	offset := []byte{}
	err := store.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == badger.ErrKeyNotFound {
			return nil
		}
		if err != nil {
			return err
		}
		offset, err = item.ValueCopy(nil)
		return err
	})

	return offset, err
}

func (store *OffsetStore) Set(offset, value []byte) error {
	return store.db.Update(func(txn *badger.Txn) error {
		return txn.Set(offset, value)
	})
}

func (store *OffsetStore) ListKeys() ([][]byte, error) {
	keys := [][]byte{}
	err := store.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.IteratorOptions{
			PrefetchValues: false,
		})
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			keys = append(keys, it.Item().KeyCopy(nil))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return keys, nil
}
