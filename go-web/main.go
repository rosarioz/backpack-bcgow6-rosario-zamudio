package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default() 

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ 
			"message": "Hola Rosario :) \n",
		})
	})

	router.GET("/transactions", GetAll)


	router.Run()

}
//estructura
type Transaction struct {
	Id 				int		`form:"id"`
	TransactionCode string	`form:"code"`
	Coin 			string	`form:"coin"`
	Amount 			float64	`form:"amount"`
	IssuingCompany 	string	`form:"company"`
	Receiver 		string	`form:"receiver"`
	TransactionDate string	`form:"date"`
}

type Transactions []Transaction

func GetAll(c *gin.Context) {
	transactions := []Transaction{}
	data, err := ioutil.ReadFile("transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	}
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"status": "ok",
			"data":   transactions,
		})
	}
}

