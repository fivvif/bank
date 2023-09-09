package repository

import (
	"bank"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateCustomer(customer bank.Customer) (int, error)
	GetCustomer(phone, password string) (bank.Customer, error)
}

type Transaction interface {
	DepositMoney(id, value int) (int, error)
	WithdrawMoney(id, value int) (int, error)
}

type Repository struct {
	Authorization
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Transaction:   NewTransPostgres(db),
	}
}
