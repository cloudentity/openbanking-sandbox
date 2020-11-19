package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
)

func (s *Server) GetBalances() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		jsonFile, err := os.Open("balances.json")
		if err != nil {
			fmt.Println(err)
		}
		defer jsonFile.Close()

		byteValue, _ := ioutil.ReadAll(jsonFile)

		var balances map[string]interface{}
		json.Unmarshal(byteValue, &balances)
		c.JSON(200, balances)
	}
}
