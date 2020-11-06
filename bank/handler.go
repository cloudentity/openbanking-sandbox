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
