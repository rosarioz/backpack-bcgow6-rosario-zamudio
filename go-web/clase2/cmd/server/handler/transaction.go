package handler

import (
	"net/http"
	

	"github.com/gin-gonic/gin"
	"github.com/rosarioz/backpack-bcgow6-rosario-zamudio/go-web/clase2/internal/transactions"
)

type Request struct{
	TransactionCode string	`json:"code"`
	Coin 			string	`json:"coin"`
	Amount 			float64	`json:"amount"`
	IssuingCompany 	string	`json:"company"`
	Receiver 		string	`json:"receiver"`
	TransactionDate string	`json:"date"`
}

type Transaction struct{
	service transactions.Service
}

func NewTransaction(t transactions.Service) *Transaction{
	return &Transaction{service: s}
}

func (t *Transaction) GetAll() gin.HandlerFunc{
	return func(c *gin.Context) {
		t, err := t.service.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if len(t) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "no hay transacciones registrados"})
			return
		}

		c.JSON(http.StatusOK, t)
	}

}

func (t *Transaction) Store() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "123456" || token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token invalido"})
			return
		}

		var req Request
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK,t)
	}
}
