package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default() 

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{ 
			"message": "Hola Rosario :) \n",
		})
	})

	data, err := os.ReadFile("./Transactions.json")
	if err != nil {
		panic(err)
	}


	json.Unmarshal(data, &transactions)

	rTransaction := router.Group("/transactions")
	{
		rTransaction.GET("/GetAll", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": transactions,
			})
		})

		rTransaction.GET(":id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				panic(err)
			}

			var filtredT []Transaction

			for i := range transactions {
				if transactions[i].Id == id {
					filtredT = append( filtredT, transactions[i])
				}
			}

			if filtredT != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"content": filtredT,
				})
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "Transaction not found",
				})
			}
		})
	}

	rTransaction.POST("", Store())

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

var transactions []Transaction

func Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req Transaction
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		req.Id = len(transactions) + 1
		transactions = append(transactions, req)
		c.JSON(http.StatusOK, req)
		file, _ := json.MarshalIndent(transactions, "", " ")
 
	_ = ioutil.WriteFile("./transactions.json", file, 0644)
	}
}








