package service

import (
	payment "github.com/elizanotbeth/test_const"
	"github.com/elizanotbeth/test_const/pckg/repository"
)

type Payment interface {
	Create(item payment.Payment) (int, error)
	GetPaymentByIdUser(user_id int) ([]payment.Payment, error)
	GetPaymentByEmail(email string) ([]payment.Payment, error)
	Update(id int, status int) error
}

type Service struct {
	Payment
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment),
	}
}
