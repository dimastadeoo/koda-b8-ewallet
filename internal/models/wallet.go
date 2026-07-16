package models

import (
	"time"
	"context"
	"github.com/jackc/pgx/v5"
)

type Wallet struct {
	ID        int64
	UserID    int64
	Balance   float64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type WalletRepository struct {
	tx pgx.Tx
}

func NewWalletRepository(tx pgx.Tx) *WalletRepository {
	return &WalletRepository{
		tx: tx,
	}
}

func (r *WalletRepository) Create(userID int64) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		INSERT INTO wallets
		(
			id_user,
			balance,
			currency,
			status
		)
		VALUES
		(
			$1,0,'IDR','active'
		)
		`,
		userID,
	)

	return err
}

func (r *WalletRepository) FindByUserID(userID int64) (*Wallet, error) {

	var wallet Wallet

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			id_user,
			balance,
			currency,
			status,
			created_at,
			updated_at
		FROM wallets
		WHERE id_user = $1
		`,
		userID,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) FindByID(id int64) (*Wallet, error) {

	var wallet Wallet

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			id_user,
			balance,
			currency,
			status,
			created_at,
			updated_at
		FROM wallets
		WHERE id = $1
		`,
		id,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}

func (r *WalletRepository) UpdateBalance(walletID int64, balance float64,) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		UPDATE wallets
		SET
			balance = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		balance,
		walletID,
	)

	return err
}

func (r *WalletRepository) UpdateStatus(walletID int64, status string,) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		UPDATE wallets
		SET
			status = $1,
			updated_at = NOW()
		WHERE id = $2
		`,
		status,
		walletID,
	)

	return err
}

func (r *WalletRepository) FindByIDForUpdate(id int64) (*Wallet, error) {

	var wallet Wallet

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			id_user,
			balance,
			currency,
			status,
			created_at,
			updated_at
		FROM wallets
		WHERE id = $1
		FOR UPDATE
		`,
		id,
	).Scan(
		&wallet.ID,
		&wallet.UserID,
		&wallet.Balance,
		&wallet.Currency,
		&wallet.Status,
		&wallet.CreatedAt,
		&wallet.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}