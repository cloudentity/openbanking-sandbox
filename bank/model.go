package main

import "time"

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
