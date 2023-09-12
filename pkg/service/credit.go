package service

import "bank/pkg/repository"

type CreditService struct {
	repo *repository.CreditPostgres
}

func NewCreditService(repo *repository.CreditPostgres) *CreditService {
	return &CreditService{repo: repo}
}
