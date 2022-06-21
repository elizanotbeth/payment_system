package repository

import (
	payment "github.com/elizanotbeth/test_const"
	"github.com/jmoiron/sqlx"
)

type Payment interface {
	Create(item payment.Payment) (int, error)
	GetPaymentByIdUser(user_id int) ([]payment.Payment, error)
	GetPaymentByEmail(email string) ([]payment.Payment, error)
	Update(id int, status int) error
}

type Repository struct {
	Payment
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
	}
}
