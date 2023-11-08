package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.POST("/add-new-user-automated-invest", addSalaryCredited)
	router.GET("/customer-information/:id", getSingleCustomerInformation)
	router.PATCH("/customer-information-update/:id", updateSingleCustomerInformation)
	router.GET("/customer-information-fund-status-check/:id", getFundStatusCheck)
	router.GET("/customer-information-all/:unique_id", getAllInformationViaUniqueId)
	// Database APIs
	router.GET("/get-all-users", getCustomerInformation)
	router.Run("localhost:9090")
}
