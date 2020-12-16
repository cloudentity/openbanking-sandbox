package main

import (
	"encoding/json"
	"fmt"
	"github.com/cloudentity/openbanking-sandbox/models"
	"github.com/pkg/errors"
	"io/ioutil"
)

type Data struct {
	Accounts     []models.OBAccount6                      `json:"accounts"`
	Balances     []models.OBReadBalance1DataBalanceItems0 `json:"balances"`
	Transactions []models.OBTransaction6                  `json:"transactions"`
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

func (a *Storage) GetAccounts(sub string) ([]models.OBAccount6, error) {
	var (
		accounts []models.OBAccount6
		data     Data
		ok       bool
	)

	if data, ok = a.userToData[sub]; !ok {
		return accounts, fmt.Errorf("no data found for user: %s", sub)
	}

	return data.Accounts, nil
}

func (a *Storage) GetBalances(sub string) ([]models.OBReadBalance1DataBalanceItems0, error) {
	var (
		balances []models.OBReadBalance1DataBalanceItems0
		data     Data
		ok       bool
	)

	if data, ok = a.userToData[sub]; !ok {
		return balances, fmt.Errorf("no data found for user: %s", sub)
	}

	return data.Balances, nil
}

func (a *Storage) GetTransactions(sub string) ([]models.OBTransaction6, error) {
	var (
		transactions []models.OBTransaction6
		data         Data
		ok           bool
	)

	if data, ok = a.userToData[sub]; !ok {
		return transactions, fmt.Errorf("no data found for user: %s", sub)
	}

	return data.Transactions, nil
}
