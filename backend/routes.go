package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Database APIs
	router.GET("/get-all-users", getCustomerInformation)
	router.POST("/add-new-user-automated-invest", addSalaryCredited)
	router.GET("/customer-information/:id", getCustomerInformationById)
	router.GET("/customer-information-all/:unique_id", getAllInformationViaUniqueId)
	router.PATCH("/customer-information-update/:id", updateSingleCustomerInformation)
	router.GET("/customer-information-fund-status-check/:id", getFundStatusCheck)

	// APIs to add
	// To get the total investment done by the customer
	// Total Earned money for all the enteries based on Unique Id
	// Total Net Worth (Saving + Investment + Future Security) for the customer Based on Unique Id
	// Future Security for all the enteries based on Unique Id
	// Saving/Emergency Fund for all the enteries based on Unique Id
	// Investments (Gold, Reits, Shares, Mutual funds) for all the enteries based on Unique Id
	// Status Check for Saving and Investment should be 30% of total Earned money based on Unique Id

	// Router localhost and port
	router.Run("localhost:9090")
}
