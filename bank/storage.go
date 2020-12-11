package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Data struct {
	Account     []models.OBAccount6           `json:"accounts"`
	Balance     models.OBReadBalance1Data     `json:"balances"`
	Transaction models.OBReadDataTransaction6 `json:"transactions"`
}

type Storage struct {
	userToData map[string]Data
}

func InitStorage() (*Storage, error) {
	var (
		as  = Storage{}
		err error
	)

	if err = as.Load(); err != nil {
		return nil, err
	}

	return &as, nil
}

func (a *Storage) Load() error {
	var (
		bs  []byte
		err error
	)

	if bs, err = ioutil.ReadFile("./data/data.json"); err != nil {
		return errors.Wrapf(err, "failed to read file")
	}

	if err = json.Unmarshal(bs, &a.userToData); err != nil {
		return errors.Wrapf(err, "failed to unmarshal data")
	}

	return nil
}

func (a *Storage) GetAccount(sub string) ([]models.OBAccount6, error) {
	var (
		accounts []models.OBAccount6
		data     Data
		ok       bool
	)

	if data, ok = a.userToData[sub]; !ok {
		return accounts, fmt.Errorf("no data found for user: %s", sub)
	}

	return data.Account, nil
}
