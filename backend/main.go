package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type investmentCalculator struct {
	ID               int     `json:"id"`
	Year             string  `json:"year"`
	SalaryCredited   float64 `json:"salary_credited"`
	Saving           float64 `json:"saving"`
	MutualFund       float64 `json:"mutual_funds"`
	Reits            float64 `json:"reits"`
	IndependentShare float64 `json:"independent_share"`
	RecurringDep     float64 `json:"recurring_deposit"`
	Gold             float64 `json:"gold"`
	FutureSecurity   float64 `json:"future_security"`
	HouseGroceries   float64 `json:"house_groceries"`
	SelfExpenses     float64 `json:"self_expense"`
	UnspentMoney     float64 `json:"unspent_money"`
}

var customerOne = []investmentCalculator{
	{
		ID:               1,
		Year:             "2023",
		SalaryCredited:   100000,
		Saving:           1000,
		MutualFund:       1000,
		Reits:            1000,
		IndependentShare: 1000,
		RecurringDep:     1000,
		Gold:             1000,
		FutureSecurity:   1000,
		HouseGroceries:   1000,
		SelfExpenses:     1000,
		UnspentMoney:     2000,
	},
}

func getCustomerInformation(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, customerOne)
}

func addInvestmentInformation(context *gin.Context) {
	var newInvestment investmentCalculator

	if err := context.BindJSON(&newInvestment); err != nil {
		return
	}

	customerOne = append(customerOne, newInvestment)

	context.IndentedJSON(http.StatusCreated, newInvestment)
}

func main() {
	router := gin.Default()
	router.GET("/customer-information", getCustomerInformation)
	router.POST("/customer-information", addInvestmentInformation)
	router.Run("localhost:9090")
}
