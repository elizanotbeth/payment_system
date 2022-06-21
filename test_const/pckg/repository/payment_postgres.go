package repository

import (
	"fmt"
	"time"

	payment "github.com/elizanotbeth/test_const"
	"github.com/jmoiron/sqlx"
)

type PaymentPostgres struct {
	db *sqlx.DB
}

func NewPaymentPostgres(db *sqlx.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) Create(item payment.Payment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var iditem int
	createPaymentQuery := fmt.Sprintf("INSERT INTO transaction (id_user, email, sum, currency, dt_create, status) values ($1, $2, $3, $4, $5, $6) RETURNING id")

	row := tx.QueryRow(createPaymentQuery, item.IdUser, item.Email, item.Sum, item.Currency, time.Now(), 1)
	err = row.Scan(&iditem)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return iditem, tx.Commit()
}

func (s *PaymentPostgres) GetPaymentByIdUser(user_id int) ([]payment.Payment, error) {
	var lists []payment.Payment

	query := fmt.Sprintf("SELECT t.id, t.email, t.sum, t.currency, t.dt_create as create, t.status FROM transaction t WHERE t.id_user = $1")
	err := s.db.Select(&lists, query, user_id)

	return lists, err
}

func (s *PaymentPostgres) GetPaymentByEmail(email string) ([]payment.Payment, error) {
	var lists []payment.Payment

	query := fmt.Sprintf("SELECT t.id, t.email, t.sum, t.currency, t.dt_create as create, t.status FROM transaction t WHERE t.email = $1")
	err := s.db.Select(&lists, query, email)

	return lists, err
}

func (s *PaymentPostgres) Update(id int, status int) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	createUsersListQuery := fmt.Sprintf("UPDATE transaction SET status = $1 WHERE id = $2")
	_, err = tx.Exec(createUsersListQuery, status, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return err
}
