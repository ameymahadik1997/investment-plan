package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"reflect"
	"strconv"
	"strings"

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

func stringToInt(stringNumber string) (int, error) {
	num, err := strconv.Atoi(stringNumber)
	if err != nil {
		return 0, err
	}
	return num, nil
}

// Database API Calls

// Function GET API to get all the Customer List
func getCustomerInformation(context *gin.Context) {
	db := dbConnect()

	var users []investmentOutput
	rows, err := db.Query("SELECT id, year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money, unique_id FROM investment_output;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user investmentOutput
		err := rows.Scan(&user.ID, &user.Year, &user.Month, &user.SalaryCredited, &user.Saving, &user.MutualFund, &user.Reits, &user.IndependentShare, &user.RecurringDep, &user.Gold, &user.FutureSecurity, &user.HouseGroceries, &user.SelfExpenses, &user.UnspentMoney, &user.UniqueId)
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

	stmt, err := db.Prepare("INSERT INTO investment_output (year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money, unique_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(autoInvestmentPlan.Year, autoInvestmentPlan.Month, autoInvestmentPlan.SalaryCredited, autoInvestmentPlan.Saving, autoInvestmentPlan.MutualFund, autoInvestmentPlan.Reits, autoInvestmentPlan.IndependentShare, autoInvestmentPlan.RecurringDep, autoInvestmentPlan.Gold, autoInvestmentPlan.FutureSecurity, autoInvestmentPlan.HouseGroceries, autoInvestmentPlan.SelfExpenses, autoInvestmentPlan.UnspentMoney, autoInvestmentPlan.UniqueId)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()
	context.IndentedJSON(http.StatusCreated, gin.H{"Message": "Information Added"})
}

// Function to get the Information of the Customer From ID
func getCustomerInformationById(context *gin.Context) {
	db := dbConnect()

	var users []investmentOutput
	paramId := context.Param("id")
	query := fmt.Sprintf("SELECT * FROM investment_output WHERE id = %s;", paramId)
	var count int
	err := db.QueryRow(query).Scan(&count)
	errString := fmt.Sprintf("Error: %s", err)

	if strings.Contains(errString, "no rows in result set") {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}
	query = "SELECT id, year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money, unique_id FROM investment_output WHERE id = ?;"
	rows, err := db.Query(query, paramId)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var user investmentOutput
		err := rows.Scan(&user.ID, &user.Year, &user.Month, &user.SalaryCredited, &user.Saving, &user.MutualFund, &user.Reits, &user.IndependentShare, &user.RecurringDep, &user.Gold, &user.FutureSecurity, &user.HouseGroceries, &user.SelfExpenses, &user.UnspentMoney, &user.UniqueId)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer db.Close()

	context.IndentedJSON(http.StatusOK, users)
}

// Function to get All the Information of the Customer From there Unique Id
func getAllInformationViaUniqueId(context *gin.Context) {
	db := dbConnect()

	var users []investmentOutput
	uniqueId := context.Param("unique_id")
	query := fmt.Sprintf("SELECT * FROM investment_output WHERE unique_id = %s;", uniqueId)
	var count int
	err := db.QueryRow(query).Scan(&count)
	errString := fmt.Sprintf("Error: %s", err)

	if strings.Contains(errString, "no rows in result set") {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}

	query = "SELECT id, year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money, unique_id FROM investment_output WHERE unique_id = ?;"
	rows, err := db.Query(query, uniqueId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user investmentOutput
		err := rows.Scan(&user.ID, &user.Year, &user.Month, &user.SalaryCredited, &user.Saving, &user.MutualFund, &user.Reits, &user.IndependentShare, &user.RecurringDep, &user.Gold, &user.FutureSecurity, &user.HouseGroceries, &user.SelfExpenses, &user.UnspentMoney, &user.UniqueId)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	defer db.Close()

	context.IndentedJSON(http.StatusOK, users)
}

// Funtion to update the fields of the customer based on the ID input
func updateSingleCustomerInformation(context *gin.Context) {
	var getInfo investmentOutput
	var newSalary investmentOutput

	db := dbConnect()

	paramId := context.Param("id")
	query := fmt.Sprintf("SELECT * FROM investment_output WHERE id = %s;", paramId)
	var count int
	err := db.QueryRow(query).Scan(&count)
	errString := fmt.Sprintf("Error: %s", err)

	if strings.Contains(errString, "no rows in result set") {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}

	if err := context.BindJSON(&getInfo); err != nil {
		return
	}

	fieldToColumn := map[string]string{
		"Year":             "year",
		"Month":            "month",
		"SalaryCredited":   "salary_credited",
		"Saving":           "saving",
		"MutualFund":       "mutual_funds",
		"Reits":            "reits",
		"IndependentShare": "independent_share",
		"RecurringDep":     "recurring_deposit",
		"Gold":             "gold",
		"FutureSecurity":   "future_security",
		"HouseGroceries":   "house_groceries",
		"SelfExpenses":     "self_expense",
		"UnspentMoney":     "unspent_money",
		"UniqueID":         "unique_id",
	}

	columnsToUpdate := make([]string, 0)

	// Iterate through the fields and check for updates
	for field, column := range fieldToColumn {
		valueOne := reflect.ValueOf(getInfo).FieldByName(field)
		valueTwo := reflect.ValueOf(newSalary).FieldByName(field)

		if valueOne.IsValid() && valueTwo.IsValid() && valueOne.Interface() != valueTwo.Interface() {
			if _, ok := valueOne.Interface().(string); ok {
				columnsToUpdate = append(columnsToUpdate, fmt.Sprintf("%s = '%v'", column, valueOne.Interface()))
			} else {
				columnsToUpdate = append(columnsToUpdate, fmt.Sprintf("%s = %v", column, valueOne.Interface()))
			}
		}
	}

	if len(columnsToUpdate) == 0 {
		context.IndentedJSON(http.StatusOK, gin.H{"Message": "No updates needed"})
		return
	}

	// Construct the SQL update statement
	query = fmt.Sprintf("UPDATE investment_output SET %s WHERE id = ?", strings.Join(columnsToUpdate, ", "))
	_, err = db.Exec(query, paramId)

	if err != nil {
		log.Fatal(err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": "Failed to update information"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"Message": "Information Updated Successfully."})
}

// Function to get the statuses of the customer's funds based on their ID input Parameter
func getFundStatusCheck(context *gin.Context) {
	db := dbConnect()
	var resultArray []Dictionary
	paramId := context.Param("id")
	query := fmt.Sprintf("SELECT * FROM investment_output WHERE id = %s;", paramId)
	var count int
	err := db.QueryRow(query).Scan(&count)
	errString := fmt.Sprintf("Error: %s", err)

	if strings.Contains(errString, "no rows in result set") {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}

	query = "SELECT id, year, month, salary_credited, saving, mutual_funds, reits, independent_share, recurring_deposit, gold, future_security, house_groceries, self_expense, unspent_money, unique_id FROM investment_output WHERE id = ?;"
	getInfo, err := db.Query(query, paramId)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Information Not Present"})
		return
	}

	defer getInfo.Close()

	// Need to work around here for the looping and fetching information

	var user investmentOutput
	for getInfo.Next() {
		err = getInfo.Scan(&user.ID, &user.Year, &user.Month, &user.SalaryCredited, &user.Saving, &user.MutualFund, &user.Reits, &user.IndependentShare, &user.RecurringDep, &user.Gold, &user.FutureSecurity, &user.HouseGroceries, &user.SelfExpenses, &user.UnspentMoney, &user.UniqueId)
		if err != nil {
			log.Fatal(err)
		}

		if user.FutureSecurity < user.SalaryCredited*0.09 {
			newDict := Dictionary{"Fund": "FutureSecurity", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "FutureSecurity", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.MutualFund < user.SalaryCredited*0.10 {
			newDict := Dictionary{"Fund": "MutualFund", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "MutualFund", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.Reits < user.SalaryCredited*0.05 {
			newDict := Dictionary{"Fund": "Reits", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "Reits", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.IndependentShare < user.SalaryCredited*0.05 {
			newDict := Dictionary{"Fund": "IndepentShare", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "IndepentShare", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.Gold < user.SalaryCredited*0.02 {
			newDict := Dictionary{"Fund": "Gold", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "Gold", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.RecurringDep < user.SalaryCredited*0.05 {
			newDict := Dictionary{"Fund": "RecurringDeposite", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "RecurringDeposite", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.HouseGroceries < user.SalaryCredited*0.30 {
			newDict := Dictionary{"Fund": "HouseGroceries", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {
			newDict := Dictionary{"Fund": "HouseGroceries", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}

		if user.SelfExpenses < user.SalaryCredited*0.30 {
			newDict := Dictionary{"Fund": "SelfExpenses", "Message": "Fund Low", "Status": "Low"}
			resultArray = append(resultArray, newDict)
		} else {

			newDict := Dictionary{"Fund": "SelfExpenses", "Message": "Fund High", "Status": "High"}
			resultArray = append(resultArray, newDict)
		}
	}

	defer db.Close()
	context.IndentedJSON(http.StatusAccepted, gin.H{"Result": resultArray})
}
