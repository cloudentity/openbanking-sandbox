package main

import (
	"github.com/boltdb/bolt"
	"github.com/pkg/errors"
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
