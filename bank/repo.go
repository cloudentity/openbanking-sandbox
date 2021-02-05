package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cloudentity/openbanking-sandbox/models"
	paymentModels "github.com/cloudentity/openbanking-sandbox/openbanking/paymentinitiation/models"
	"github.com/pkg/errors"
	bolt "go.etcd.io/bbolt"
)

var (
	bucketName = []byte(`users`)
)

type UserRepo struct {
	*bolt.DB
}

type Data struct {
	Accounts     []models.OBAccount6                      `json:"accounts"`
	Balances     []models.OBReadBalance1DataBalanceItems0 `json:"balances"`
	Transactions []models.OBTransaction6                  `json:"transactions"`
	Payments     []paymentModels.OBWriteDomesticResponse5 `json:"payments"`
}

type UserToDataFile map[string]Data

type ErrNotFound struct {
	name string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("unable to find resource %s", e.name)
}

func readUserToDataFile() (UserToDataFile, error) {
	var (
		bs   []byte
		u2df UserToDataFile
		err  error
	)

	if bs, err = ioutil.ReadFile("./data/data.json"); err != nil {
		return u2df, errors.Wrapf(err, "failed to read file")
	}

	if err = json.Unmarshal(bs, &u2df); err != nil {
		return u2df, errors.Wrapf(err, "failed to unmarshal data")
	}

	return u2df, nil
}

func NewUserRepo() (UserRepo, error) {
	var (
		userRepo UserRepo
		u2df     UserToDataFile
		err      error
	)

	// create db
	if userRepo.DB, err = bolt.Open("data/tppdb.db", 0644, &bolt.Options{Timeout: 3 * time.Second}); err != nil {
		return userRepo, errors.Wrapf(err, "failed to open db")
	}

	// read init data from file
	if u2df, err = readUserToDataFile(); err != nil {
		return userRepo, errors.Wrapf(err, "failed to read data file")
	}

	// setup bucket and default data
	if err = userRepo.Update(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket
		if bucket, err = tx.CreateBucketIfNotExists(bucketName); err != nil {
			return errors.Wrapf(err, "failed to create bucket")
		}

		for k, v := range u2df {
			var (
				valBytes []byte
				err      error
			)

			if valBytes, err = json.Marshal(v); err != nil {
				return errors.Wrapf(err, "failed to unmarshal data value from file")
			}

			if err = bucket.Put([]byte(k), valBytes); err != nil {
				return errors.Wrapf(err, "failed to put value in bucket")
			}
		}

		return nil
	}); err != nil {
		return userRepo, err
	}

	return userRepo, nil
}

func (us *UserRepo) GetAccounts(sub string) ([]models.OBAccount6, error) {
	var (
		data Data
		err  error
	)

	if err = us.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err = json.Unmarshal(v, &data); err != nil {
			return errors.Wrapf(err, "failed to unmarshal json")
		}
		return nil
	}); err != nil {
		return data.Accounts, errors.Wrapf(err, "failed to read account data")
	}

	return data.Accounts, nil
}

func (us *UserRepo) GetBalances(sub string) ([]models.OBReadBalance1DataBalanceItems0, error) {
	var (
		data Data
		err  error
	)

	if err = us.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err = json.Unmarshal(v, &data); err != nil {
			return errors.Wrapf(err, "failed to unmarshal json")
		}
		return nil
	}); err != nil {
		return data.Balances, errors.Wrapf(err, "failed to read balances data")
	}

	return data.Balances, nil
}

func (us *UserRepo) GetTransactions(sub string) ([]models.OBTransaction6, error) {
	var (
		data Data
		err  error
	)

	if err = us.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err = json.Unmarshal(v, &data); err != nil {
			return errors.Wrapf(err, "failed to unmarshal json")
		}
		return nil
	}); err != nil {
		return data.Transactions, errors.Wrapf(err, "failed to read transaction data")
	}

	return data.Transactions, nil
}

// either this or the web handler will need to trigger a routine to modify accounts/transactions/balances whenever the "payment" goes through
// additionally, can just have a separate go routine running that always listens for these things
// this will create the domestic payment resource
func (ur *UserRepo) CreateDomesticPayment(sub string, payment paymentModels.OBWriteDomesticResponse5) error {
	var (
		data      Data
		err       error
		dataBytes []byte
	)

	if err = ur.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))

		if err = json.Unmarshal(v, &data); err != nil {
			return err
		}

		data.Payments = append(data.Payments, payment)

		if dataBytes, err = json.Marshal(data); err != nil {
			return err
		}

		if err = b.Put([]byte(sub), dataBytes); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return errors.Wrapf(err, "failed to update db")
	}

	return nil
}

func (ur *UserRepo) GetDomesticPayment(sub, domesticPaymentID string) (paymentModels.OBWriteDomesticResponse5, error) {
	var (
		data    Data
		payment paymentModels.OBWriteDomesticResponse5
		err     error
	)

	if err = ur.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		v := b.Get([]byte(sub))
		if err = json.Unmarshal(v, &data); err != nil {
			return errors.Wrapf(err, "failed to unmarshal data")
		}
		return nil
	}); err != nil {
		return payment, errors.Wrapf(err, "failed to read payment data")
	}

	for _, p := range data.Payments {
		if *p.Data.DomesticPaymentID == domesticPaymentID {
			return p, nil
		}
	}

	return payment, ErrNotFound{fmt.Sprintf("domestic-payment with id %s", domesticPaymentID)}
}
