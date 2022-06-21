package service

import (
	payment "github.com/elizanotbeth/test_const"
	"github.com/elizanotbeth/test_const/pckg/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(item payment.Payment) (int, error) {
	return s.repo.Create(item)
}
func (s *PaymentService) GetPaymentByIdUser(user_id int) ([]payment.Payment, error) {
	return s.repo.GetPaymentByIdUser(user_id)
}

func (s *PaymentService) GetPaymentByEmail(email string) ([]payment.Payment, error) {
	return s.repo.GetPaymentByEmail(email)
}

func (s *PaymentService) Update(id int, status int) error {
	return s.repo.Update(id, status)
}
