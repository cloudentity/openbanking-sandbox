package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

type GetAccounts struct {
	Data  Data  `json:"Data"`
	Meta  Meta  `json:"Meta"`
	Links Links `json:"Links"`
}

type Meta struct {
	TotalPages int `json:"TotalPages"`
}

type Links struct {
	// example: https://api.alphabank.com/open-banking/v3.1/aisp/accounts/
	Self string `json:"Self"`
}

type Data struct {
	Account []Account `json:"Account"`
}

type Account struct {
	// example: 22289
	AccountID string `json:"AccountId"`
	// example: Enabled
	Status string `json:"Status"`
	// example: 2019-01-01T06:06:06+00:00
	StatusUpdateDateTime time.Time `json:"StatusUpdateDateTime"`
	// example: GBP
	Currency string `json:"Currency"`
	// example: personal
	AccountType string `json:"AccountType"`
	// example: CurrentAccount
	AccountSubType string `json:"AccountSubType"`
	// example: Bills
	Nickname string           `json:"Nickname"`
	Account  []AccountDetails `json:"Account"`
}

type AccountDetails struct {
	// example: UK.OBIE.SortCodeAccountNumber
	SchemeName string `json:"SchemeName"`
	// example: 80200110203345
	Identification string `json:"Identification"`
	// example Mr Kevin
	Name string `json:"Name"`
	// example 00021
	SecondaryIdentification string `json:"SecondaryIdentification"`
}

type AccountsStorage struct {
	userToAccounts map[string][]Account
}

func InitAccountsStorage() (*AccountsStorage, error) {
	var (
		as  = AccountsStorage{}
		err error
	)

	if err = as.Load(); err != nil {
		return nil, err
	}

	return &as, nil
}

func (a *AccountsStorage) Load() error {
	var (
		bs  []byte
		err error
	)

	if bs, err = ioutil.ReadFile("./data/accounts.json"); err != nil {
		return errors.Wrapf(err, "failed to read file")
	}

	if err = json.Unmarshal(bs, &a.userToAccounts); err != nil {
		return errors.Wrapf(err, "failed to unmarshal accounts")
	}

	return nil
}

func (a *AccountsStorage) GetAccount(sub string) ([]Account, error) {
	var (
		accounts []Account
		ok       bool
	)

	if accounts, ok = a.userToAccounts[sub]; !ok {
		return accounts, fmt.Errorf("no accounts found for user: %s", sub)
	}

	return accounts, nil
}
