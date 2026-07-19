package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type User struct {
	ID            int64
	Pin           string
	Email         string
	HPNumber      string
	StatusAccount string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type UserRepository struct {
	tx pgx.Tx
}

func NewUserRepository(tx pgx.Tx) *UserRepository {
	return &UserRepository{
		tx: tx,
	}
}

func (r *UserRepository) Create(user User) (int64, error) {

	var id int64

	err := r.tx.QueryRow(
		context.Background(),
		`
		INSERT INTO users
		(
			pin,
			email,
			hp_number
		)
		VALUES
		(
			$1,$2,$3
		)
		RETURNING id
		`,
		user.Pin,
		user.Email,
		user.HPNumber,
	).Scan(&id)

	return id, err

}

func (r *UserRepository) FindByPhone(phone string) (*User, error) {

	var user User

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			pin,
			email,
			hp_number,
			status_account
		FROM users
		WHERE hp_number=$1
		`,
		phone,
	).Scan(
		&user.ID,
		&user.Pin,
		&user.Email,
		&user.HPNumber,
		&user.StatusAccount,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) FindByEmail(email string) (*User, error) {

	var user User

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			pin,
			email,
			hp_number,
			status_account
		FROM users
		WHERE email=$1
		`,
		email,
	).Scan(
		&user.ID,
		&user.Pin,
		&user.Email,
		&user.HPNumber,
		&user.StatusAccount,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *UserRepository) FindByEmailAndPhone(email string, hp string) (*User, error) {

	var user User

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			email,
			hp_number,
			pin,
			status_account,
			created_at,
			updated_at
		FROM users
		WHERE email=$1
		AND hp_number=$2
		`,
		email,
		hp,
	).Scan(
		&user.ID,
		&user.Email,
		&user.HPNumber,
		&user.Pin,
		&user.StatusAccount,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindByID(id int64) (*User, error) {
	var user User

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			pin,
			email,
			hp_number,
			status_account
		FROM users
		WHERE id = $1
		`,
		id,
	).Scan(
		&user.ID,
		&user.Pin,
		&user.Email,
		&user.HPNumber,
		&user.StatusAccount,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdatePin(id int64, newPin string) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		UPDATE users
		SET
			pin=$1,
			updated_at=NOW()
		WHERE id=$2
		`,
		newPin,
		id,
	)

	return err

}

func (r *UserRepository) UpdateStatus(id int64, status string) error {
	_, err := r.tx.Exec(
		context.Background(),
		`
		UPDATE users
		SET
			status_account=$1,
			updated_at=NOW()
		WHERE id=$2
		`,
		status,
		id,
	)

	return err

}
