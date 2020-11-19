package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func (s *Server) GetTransactions() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jsonFile, err := os.Open("transactions.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var transactions map[string]interface{}
		json.Unmarshal(byteValue, &transactions)
		c.JSON(200, transactions)
	}
}
