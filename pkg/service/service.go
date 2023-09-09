<<<<<<< HEAD
package service

import "bank/pkg/repository"

type Authorization interface {
}
type Transaction interface {
}

type Service struct {
	Authorization
	Transaction
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}

}
=======
package service

import (
	"bank"
	"bank/pkg/repository"
)

type Authorization interface {
	CreateCustomer(customer bank.Customer) (int, error)
	GenerateToken(phone, password string) (string, error)
	ParseToken(token string) (int, error)
}
type Transaction interface {
}

type Service struct {
	Authorization
	Transaction
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}

}
>>>>>>> f48611f (first commit)
