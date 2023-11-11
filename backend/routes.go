package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/customer-information-fund-status-check/:id", getFundStatusCheck)
	// Database APIs
	router.GET("/get-all-users", getCustomerInformation)
	router.POST("/add-new-user-automated-invest", addSalaryCredited)
	router.GET("/customer-information/:id", getCustomerInformationById)
	router.GET("/customer-information-all/:unique_id", getAllInformationViaUniqueId)
	router.PATCH("/customer-information-update/:id", updateSingleCustomerInformation)
	router.Run("localhost:9090")
}
