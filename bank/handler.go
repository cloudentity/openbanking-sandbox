package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		// todo call ob introspection

		// ReadAccountsBasic - bez AccountDetails
		// ReadAccountsDetail - wszystko
		// ReadAccountsDetail - GET by ID

		// mocked data
		response := GetAccounts{
			Data: Data{
				Account: []Account{
					{
						AccountID:      "22289",
						Status:         "Enabled",
						Currency:       "GBP",
						Nickname:       "Bills",
						AccountType:    "Personal",
						AccountSubType: "CurrentAccount",
						Account: []AccountDetails{
							{
								SchemeName:              "UK.OBIE.SortCodeAccountNumber",
								Identification:          "80200110203345",
								Name:                    "Mr Kevin",
								SecondaryIdentification: "00021",
							},
						},
					},
					{
						AccountID:      "31820",
						Status:         "Enabled",
						Currency:       "GBP",
						Nickname:       "Household",
						AccountType:    "Personal",
						AccountSubType: "CurrentAccount",
						Account: []AccountDetails{
							{
								SchemeName:     "UK.OBIE.SortCodeAccountNumber",
								Identification: "80200110203348",
								Name:           "Mr Kevin",
							},
						},
					},
				},
			},
		}
		response.Meta.TotalPages = len(response.Data.Account)
		response.Links.Self = fmt.Sprintf("http://localhost:%s/accounts", strconv.Itoa(s.Config.Port))

		c.PureJSON(http.StatusOK, response)
	}
}

type InternalAccounts struct {
	Accounts []InternalAccount `json:"accounts"`
}

type InternalAccount struct {
	ID   string `json:"id"`
	Name string `json:"Name"`
}

// this API is bank specific. It should return all users's account.
func (s *Server) InternalGetAccounts() func(*gin.Context) {
	return func(c *gin.Context) {
		response := InternalAccounts{
			Accounts: []InternalAccount{
				{
					ID:   "27 1140 2004 0000 3002 0135 5387",
					Name: "Bills",
				},
				{
					ID:   "17 2240 1401 0000 000 0155 1312",
					Name: "Household",
				},
			},
		}

		c.PureJSON(http.StatusOK, response)
	}
}
