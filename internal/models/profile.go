package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	ID         int64
	UserID     int64
	NIK        *string
	Name       string
	Address    *string
	Gender     string
	PlaceBirth *string
	DateBirth  *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type ProfileRepository struct {
	tx pgx.Tx
}

func NewProfileRepository(tx pgx.Tx) *ProfileRepository {
	return &ProfileRepository{
		tx: tx,
	}
}

func (r *ProfileRepository) Create(profile Profile) (int64, error) {

	var id int64

	err := r.tx.QueryRow(
		context.Background(),
		`
		INSERT INTO profiles
		(
			id_user,
			nik,
			name,
			address,
			gender,
			place_birth,
			date_birth
		)
		VALUES
		(
			$1,$2,$3,$4,$5,$6,$7
		)
		RETURNING id
		`,
		profile.UserID,
		profile.NIK,
		profile.Name,
		profile.Address,
		profile.Gender,
		profile.PlaceBirth,
		profile.DateBirth,
	).Scan(&id)

	return id, err
}

func (r *ProfileRepository) FindByUserID(userID int64) (*Profile, error) {

	var profile Profile

	err := r.tx.QueryRow(
		context.Background(),
		`
		SELECT
			id,
			id_user,
			nik,
			name,
			address,
			gender,
			place_birth,
			date_birth,
			created_at,
			updated_at
		FROM profiles
		WHERE id_user = $1
		`,
		userID,
	).Scan(
		&profile.ID,
		&profile.UserID,
		&profile.NIK,
		&profile.Name,
		&profile.Address,
		&profile.Gender,
		&profile.PlaceBirth,
		&profile.DateBirth,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (r *ProfileRepository) Update(profile Profile) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		UPDATE profiles
		SET
			nik = $1,
			name = $2,
			address = $3,
			gender = $4,
			place_birth = $5,
			date_birth = $6,
			updated_at = NOW()
		WHERE id_user = $7
		`,
		profile.NIK,
		profile.Name,
		profile.Address,
		profile.Gender,
		profile.PlaceBirth,
		profile.DateBirth,
		profile.UserID,
	)

	return err
}

func (r *ProfileRepository) Delete(userID int64) error {

	_, err := r.tx.Exec(
		context.Background(),
		`
		DELETE FROM profiles
		WHERE id_user = $1
		`,
		userID,
	)

	return err
}
