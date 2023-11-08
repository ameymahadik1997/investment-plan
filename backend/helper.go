package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Function to connect to the DB
func dbConnect() *sql.DB {
	// MySQL connection
	viper.SetConfigFile("keys.json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	dbUsername := viper.GetString("DB_USERNAME")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbHost := viper.GetString("DB_HOST")
	dbPort := viper.GetString("DB_PORT")
	dbName := viper.GetString("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("Error connecting to database")
	}

	return db
}

var customerOne = []investmentOutput{}

func getSingleCustomerInformationById(id string) (*investmentOutput, error) {
	idConvert, _ := stringToInt(id)
	for index, investmentCalculatorId := range customerOne {
		if investmentCalculatorId.ID == idConvert {
			return &customerOne[index], nil
		}
	}
	return nil, errors.New("Investment Information not found!")
}

func getCustomerInformationByUniqueId(id string) ([]*investmentOutput, error) {
	var resultArray []*investmentOutput // Change the type of resultArray to []*investmentOutput
	uniqueIdConvert, _ := stringToInt(id)
	for index, investmentCalculatorId := range customerOne {
		if int(investmentCalculatorId.UniqueId) == uniqueIdConvert {
			resultArray = append(resultArray, &customerOne[index])
		}
	}
	if len(resultArray) == 0 {
		return nil, errors.New("Investment Information not found!")
	}
	return resultArray, nil
}

func getSingleCustomerInformation(context *gin.Context) {
	id := context.Param("id")
	getInfo, err := getSingleCustomerInformationById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Information was not found!"})
		return
	}
	context.IndentedJSON(http.StatusOK, getInfo)
}

func stringToInt(stringNumber string) (int, error) {
	num, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func updateSingleCustomerInformation(context *gin.Context) {
	var newSalary investmentOutput

	id := context.Param("id")
	getInfo, err := getSingleCustomerInformationById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Information was not found!"})
		return
	}

	if err := context.BindJSON(&getInfo); err != nil {
		return
	}

	newSalary.ID = getInfo.ID
	newSalary.Year = getInfo.Year
	newSalary.SalaryCredited = getInfo.SalaryCredited
	newSalary.Month = getInfo.Month
	newSalary.Saving = getInfo.Saving
	newSalary.MutualFund = getInfo.MutualFund
	newSalary.Reits = getInfo.Reits
	newSalary.IndependentShare = getInfo.IndependentShare
	newSalary.Gold = getInfo.Gold
	newSalary.RecurringDep = getInfo.RecurringDep
	newSalary.FutureSecurity = getInfo.FutureSecurity
	newSalary.HouseGroceries = getInfo.HouseGroceries
	newSalary.SelfExpenses = getInfo.SelfExpenses
	newSalary.UnspentMoney = getInfo.UnspentMoney
	newSalary.UniqueId = getInfo.UniqueId

	context.IndentedJSON(http.StatusCreated, newSalary)
}

func getFundStatusCheck(context *gin.Context) {
	var resultArray = []Dictionary{}
	id := context.Param("id")
	getInfo, err := getSingleCustomerInformationById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Information was not found!"})
		return
	}

	if getInfo.FutureSecurity < getInfo.SalaryCredited*0.09 {
		newDict := Dictionary{"Fund": "FutureSecurity", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "FutureSecurity", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.MutualFund < getInfo.SalaryCredited*0.10 {
		newDict := Dictionary{"Fund": "MutualFund", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "MutualFund", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.Reits < getInfo.SalaryCredited*0.05 {
		newDict := Dictionary{"Fund": "Reits", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "Reits", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.IndependentShare < getInfo.SalaryCredited*0.05 {
		newDict := Dictionary{"Fund": "IndepentShare", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "IndepentShare", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.Gold < getInfo.SalaryCredited*0.02 {
		newDict := Dictionary{"Fund": "Gold", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "Gold", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.RecurringDep < getInfo.SalaryCredited*0.05 {
		newDict := Dictionary{"Fund": "RecurringDeposite", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "RecurringDeposite", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.HouseGroceries < getInfo.SalaryCredited*0.30 {
		newDict := Dictionary{"Fund": "HouseGroceries", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {
		newDict := Dictionary{"Fund": "HouseGroceries", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	if getInfo.SelfExpenses < getInfo.SalaryCredited*0.30 {
		newDict := Dictionary{"Fund": "SelfExpenses", "Message": "Fund Low", "Status": "Low"}
		resultArray = append(resultArray, newDict)
	} else {

		newDict := Dictionary{"Fund": "SelfExpenses", "Message": "Fund High", "Status": "High"}
		resultArray = append(resultArray, newDict)
	}

	context.IndentedJSON(http.StatusAccepted, gin.H{"Result": resultArray})
}

func getAllInformationViaUniqueId(context *gin.Context) {
	uniqueId := context.Param("unique_id")
	getInfo, err := getCustomerInformationByUniqueId(uniqueId)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Information was not found!"})
		return
	}
	context.IndentedJSON(http.StatusOK, getInfo)
}

// Database API Calls

// Function GET API to get all the Customer List
func getCustomerInformation(context *gin.Context) {
	db := dbConnect()

	var users []investmentOutput
	rows, err := db.Query("SELECT id, year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money FROM investment_output;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user investmentOutput
		err := rows.Scan(&user.ID, &user.Year, &user.Month, &user.SalaryCredited, &user.Saving, &user.MutualFund, &user.Reits, &user.IndependentShare, &user.RecurringDep, &user.Gold, &user.FutureSecurity, &user.HouseGroceries, &user.SelfExpenses, &user.UnspentMoney)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer db.Close()

	context.IndentedJSON(http.StatusOK, users)
}

// Function POST API to add the data for customer automatedly
func addSalaryCredited(context *gin.Context) {
	var newSalary investmentAutoCalculator
	var autoInvestmentPlan investmentOutput

	if err := context.BindJSON(&newSalary); err != nil {
		return
	}

	salaryCredited := newSalary.SalaryCredited
	saving := newSalary.SalaryCredited * 0.20
	mutualFund := newSalary.SalaryCredited * 0.10
	reits := newSalary.SalaryCredited * 0.05
	share := newSalary.SalaryCredited * 0.05
	gold := newSalary.SalaryCredited * 0.02
	recurringDeposit := newSalary.SalaryCredited * 0.05
	futureSecurity := newSalary.SalaryCredited * 0.09
	houseGroceries := newSalary.SalaryCredited * 0.30
	selfExpenses := newSalary.SalaryCredited * 0.20

	// Logging the Outputs
	fmt.Printf("SalaryCredited: %v\n", salaryCredited)
	fmt.Printf("Saving: %v\n", saving)
	fmt.Printf("MutualFund: %v\n", mutualFund)
	fmt.Printf("Reits: %v\n", reits)
	fmt.Printf("Share: %v\n", share)
	fmt.Printf("Gold: %v\n", gold)
	fmt.Printf("RecurringDeposit: %v\n", recurringDeposit)
	fmt.Printf("FutureSecurity: %v\n", futureSecurity)
	fmt.Printf("HouseGroceries: %v\n", houseGroceries)
	fmt.Printf("SelfExpenses: %v\n", selfExpenses)

	db := dbConnect()

	autoInvestmentPlan.ID = newSalary.ID
	autoInvestmentPlan.Year = newSalary.Year
	autoInvestmentPlan.SalaryCredited = newSalary.SalaryCredited
	autoInvestmentPlan.Month = newSalary.Month
	autoInvestmentPlan.Saving = saving
	autoInvestmentPlan.MutualFund = mutualFund
	autoInvestmentPlan.Reits = reits
	autoInvestmentPlan.IndependentShare = share
	autoInvestmentPlan.Gold = gold
	autoInvestmentPlan.RecurringDep = recurringDeposit
	autoInvestmentPlan.FutureSecurity = futureSecurity
	autoInvestmentPlan.HouseGroceries = houseGroceries
	autoInvestmentPlan.SelfExpenses = selfExpenses
	if newSalary.UniqueId != 0 {
		autoInvestmentPlan.UniqueId = newSalary.UniqueId

	} else {
		autoInvestmentPlan.UniqueId = rand.Int31()
	}

	stmt, err := db.Prepare("INSERT INTO investment_output (year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(autoInvestmentPlan.Year, autoInvestmentPlan.Month, autoInvestmentPlan.SalaryCredited, autoInvestmentPlan.Saving, autoInvestmentPlan.MutualFund, autoInvestmentPlan.Reits, autoInvestmentPlan.IndependentShare, autoInvestmentPlan.RecurringDep, autoInvestmentPlan.Gold, autoInvestmentPlan.FutureSecurity, autoInvestmentPlan.HouseGroceries, autoInvestmentPlan.SelfExpenses, autoInvestmentPlan.UnspentMoney)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	context.IndentedJSON(http.StatusCreated, gin.H{"Message": "Information Added"})
}
