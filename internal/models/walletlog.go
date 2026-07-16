package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type WalletLog struct {
	ID          int64
	WalletID    int64
	Action      string
	Description string
	CreatedAt   time.Time
}

type WalletLogRepository struct {
	tx pgx.Tx
}

func NewWalletLogRepository(tx pgx.Tx) *WalletLogRepository {
	return &WalletLogRepository{
		tx: tx,
	}
}

func (r *WalletLogRepository) Create(
	log WalletLog,
) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		INSERT INTO wallet_logs
		(
			wallet_id,
			action,
			description
		)
		VALUES
		(
			$1,$2,$3
		)
		`,
		log.WalletID,
		log.Action,
		log.Description,
	)

	return err
}