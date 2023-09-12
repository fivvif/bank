package bank

import "time"

type Credit struct {
	Id         int
	CustomerId int
	Value      int
	Percentage int
	LoanPeriod time.Time
}

// ОБЯЗАТЕЛЬНЫЕ ПЛАТЕЖИ