package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type UserLogRepository struct {
	tx pgx.Tx
}

func NewUserLogRepository(tx pgx.Tx) *UserLogRepository {
	return &UserLogRepository{
		tx: tx,
	}
}

func (r *UserLogRepository) Create(
	userID int64,
	sessionID int64,
	activity string,
	ip string,
) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		INSERT INTO users_logs
		(
			id_user,
			id_sesion,
			activity_detail,
			ip_address
		)
		VALUES
		(
			$1,$2,$3,$4
		)
		`,
		userID,
		sessionID,
		activity,
		ip,
	)

	return err
}
