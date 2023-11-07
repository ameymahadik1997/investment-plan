package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	router.GET("/customer-information", getCustomerInformation)
	router.POST("/customer-information", addInvestmentInformation)
	router.POST("/customer-information-salary", addSalaryCredited)
	router.GET("/customer-information/:id", getSingleCustomerInformation)
	router.PATCH("/customer-information-update/:id", updateSingleCustomerInformation)
	router.GET("/customer-information-fund-status-check/:id", getFundStatusCheck)
	router.Run("localhost:9090")
}
