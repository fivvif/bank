package repository

import "github.com/jmoiron/sqlx"

type CreditPostgres struct {
	db *sqlx.DB
}

func NewCreditPostgres(db *sqlx.DB) *CreditPostgres {
	return &CreditPostgres{db: db}
}
