package main

type investmentAutoCalculator struct {
	ID             int     `json:"id"`
	Year           string  `json:"year"`
	Month          string  `json:"month"`
	SalaryCredited float64 `json:"salary_credited"`
}

type investmentOutput struct {
	ID               int     `json:"id"`
	Year             string  `json:"year"`
	Month            string  `json:"month"`
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
	UniqueId         int32   `json:"unique_id"`
}

type Dictionary map[string]interface{}
