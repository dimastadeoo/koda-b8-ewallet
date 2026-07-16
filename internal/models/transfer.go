package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Transfer struct {
	ID               int64
	TransactionID    int64
	SenderWalletID   int64
	ReceiverWalletID int64
	Notes            string
	CreatedAt        time.Time
}

type TransferHistory struct {
	ID         int64
	Name       string
	Phone      string
	Amount     float64
	Status     string
	Notes      string
	CreatedAt  time.Time
}

type TransferRepository struct {
	tx pgx.Tx
}

func NewTransferRepository(tx pgx.Tx) *TransferRepository {
	return &TransferRepository{
		tx: tx,
	}
}

func (r *TransferRepository) Create(
	transfer Transfer,
) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		INSERT INTO transfers
		(
			trx_id,
			sender_wallet_id,
			receiver_wallet_id,
			notes
		)
		VALUES
		(
			$1,$2,$3,$4
		)
		`,
		transfer.TransactionID,
		transfer.SenderWalletID,
		transfer.ReceiverWalletID,
		transfer.Notes,
	)

	return err
}

func (r *TransferRepository) History(
	walletID int64,
) ([]TransferHistory, error) {

	rows, err := r.tx.Query(
		context.Background(),
		`
		SELECT
			tf.id,
			p.name,
			u.hp_number,
			t.amount,
			t.status,
			tf.notes,
			tf.created_at
		FROM transfers tf

		JOIN transactions t
			ON t.id = tf.trx_id

		JOIN wallets rw
			ON rw.id = tf.receiver_wallet_id

		JOIN users u
			ON u.id = rw.id_user

		JOIN profiles p
			ON p.id_user = u.id

		WHERE tf.sender_wallet_id = $1

		ORDER BY tf.created_at DESC
		`,
		walletID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var histories []TransferHistory

	for rows.Next() {

		var h TransferHistory

		err := rows.Scan(
			&h.ID,
			&h.Name,
			&h.Phone,
			&h.Amount,
			&h.Status,
			&h.Notes,
			&h.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		histories = append(histories, h)
	}

	return histories, rows.Err()
}