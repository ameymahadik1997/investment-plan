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
	router.GET("/get-total-investment/:unique_id", getTotalInvestmentByCustomer)
	router.GET("/get-total-money-earned/:unique_id", getTotalEarnedMoneyByCustomer)
	router.GET("/get-total-net-worth/:unique_id", getTotalNetWorthByCustomer)
	router.GET("/get-total-future-securities/:unique_id", getTotalFutureSecuritiesByCustomer)
	router.GET("/get-total-emergency-liquid-fund/:unique_id", getTotalEmergencyFundByCustomer)

	// APIs to add

	// Investments (Gold, Reits, Shares, Mutual funds) for all the enteries based on Unique Id
	// Status Check for Saving and Investment should be 30% of total Earned money based on Unique Id
	// Add a table to add the default percentage ratio for the auto fund addition for the new entries
	// Add an API to get the current percentage for the customer based on Unique ID
	// Add an API to add the new percentage allocation for the customer
	// Add an assertion for the max or min limit for any of the specific fund allocation percentage

	// Router localhost and port
	router.Run("localhost:9090")
}
