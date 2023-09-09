<<<<<<< HEAD
package repository

type Authorization interface {
}

type Transaction interface {
}

type Repository struct {
	Authorization
	Transaction
}

func NewRepository() *Repository {
	return &Repository{}
}
=======
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
}

type Repository struct {
	Authorization
	Transaction
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
>>>>>>> f48611f (first commit)
