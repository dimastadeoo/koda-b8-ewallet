package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	ID              int64
	WalletID        int64
	TransactionType string
	Amount          float64
	Status          string
	Description     string
	CreatedAt       time.Time
	CompletedAt     *time.Time
}

type TransactionRepository struct {
	tx pgx.Tx
}

func NewTransactionRepository(tx pgx.Tx) *TransactionRepository {
	return &TransactionRepository{
		tx: tx,
	}
}

func (r *TransactionRepository) Create(
	transaction Transaction,
) (int64, error) {

	var id int64

	err := r.tx.QueryRow(
		context.Background(),
		`
		INSERT INTO transactions
		(
			id_wallet,
			transaction_type,
			amount,
			status,
			description,
			completed_at
		)
		VALUES
		(
			$1,$2,$3,$4,$5,NOW()
		)
		RETURNING id
		`,
		transaction.WalletID,
		transaction.TransactionType,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
	).Scan(&id)

	return id, err
}

// func (r *TransactionRepository) GetByWalletID(
// 	walletID int64,
// ) ([]Transaction, error)