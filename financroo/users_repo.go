package main

import (
	"encoding/json"


	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

type User struct {
	Sub   string          `json:"sub"`
	Banks []ConnectedBank `json:"banks"`
}

type ConnectedBank struct {
	BankID   string `json:"bank_id"`
	IntentID string `json:"intent_id"`
}

type UserRepo struct {
	*bolt.DB
}

var bucket = []byte("users")

func NewUserRepo(db *bolt.DB) (UserRepo, error) {
	var (
		repo = UserRepo{}
		err  error
	)

	if err = db.Update(func(tx *bolt.Tx) error {
		if _, err = tx.CreateBucketIfNotExists(bucket); err != nil {
			return errors.Wrapf(err, "failed to create bucket")
		}

		return nil
	}); err != nil {
		return repo, err
	}

	repo.DB = db

	return repo, nil
}

func (u *UserRepo) Get(sub string) (User, error) {
	var (
		user User
		bs   []byte
		err  error
	)

	if err = u.DB.View(func(tx *bolt.Tx) error {
		bs = tx.Bucket(bucket).Get([]byte(sub))

		return nil
	}); err != nil {
		return user, errors.Wrapf(err, "failed to get user")
	}

	if err = json.Unmarshal(bs, &user); err != nil {
		return user, errors.Wrapf(err, "failed to unmarshal user")
	}

	return user, nil
}

func (u *UserRepo) Set(user User) error {
	var (
		bs  []byte
		err error
	)

	if bs, err = json.Marshal(&user); err != nil {
		return errors.Wrapf(err, "failed to marshal user")
	}

	if err = u.DB.Update(func(tx *bolt.Tx) error {
		return tx.Bucket(bucket).Put([]byte(user.Sub), bs)
	}); err != nil {
		return errors.Wrapf(err, "failed to update user")
	}

	return nil
}
