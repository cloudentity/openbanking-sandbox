package main

import (
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
	"time"
)

func InitDB() (*bolt.DB, error) {
	var (
		db  *bolt.DB
		err error
	)

	if db, err = bolt.Open("./my.db", 0600, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return nil, errors.Wrapf(err, "failed to open db")
	}

	return db, nil
}
