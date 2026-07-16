package models

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/jackc/pgx/v5"
)

type Session struct {
	ID        int64
	Token     string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginSession struct {
	Session Session
	User    User
	Profile Profile
	Wallet  Wallet
}

type SessionRepository struct {
	tx pgx.Tx
}

func NewSessionRepository(tx pgx.Tx) *SessionRepository {
	return &SessionRepository{
		tx: tx,
	}
}

func generateToken() (string, error) {

	bytes := make([]byte, 32)

	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}

func (r *SessionRepository) Create(status string) (*Session, error) {

	token, err := generateToken()
	if err != nil {
		return nil, err
	}

	session := Session{}

	err = r.tx.QueryRow(
		context.Background(),
		`
		INSERT INTO sessions
		(
			token,
			status
		)
		VALUES
		(
			$1,$2
		)
		RETURNING
			id,
			token,
			status,
			created_at,
			updated_at
		`,
		token,
		status,
	).Scan(
		&session.ID,
		&session.Token,
		&session.Status,
		&session.CreatedAt,
		&session.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &session, nil
}
