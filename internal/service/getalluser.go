package service

import (
	"context"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	conn *pgx.Conn
}

func NewUserService(conn *pgx.Conn) *UserService {
	return &UserService{
		conn: conn,
	}
}

func (s *UserService) GetAll() ([]models.UserList, error) {

	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	userRepo := models.NewUserRepository(tx)

	users, err := userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
