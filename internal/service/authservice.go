package service

import (
	"context"
	"errors"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	conn *pgx.Conn
}

func NewAuthService(conn *pgx.Conn) *AuthService {
	return &AuthService{
		conn: conn,
	}
}

func (s *AuthService) Register(
	email string,
	hp string,
	pin string,
	name string,
	gender string,
) error {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	userRepo := models.NewUserRepository(tx)
	profileRepo := models.NewProfileRepository(tx)
	walletRepo := models.NewWalletRepository(tx)

	// Hash PIN
	hashPin, err := bcrypt.GenerateFromPassword(
		[]byte(pin),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	// Insert User
	userID, err := userRepo.Create(models.User{
		Email:    email,
		HPNumber: hp,
		Pin:      string(hashPin),
	})

	if err != nil {
		return err
	}

	// Insert Profile
	_, err = profileRepo.Create(models.Profile{
		UserID: userID,
		Name:   name,
		Gender: gender,
	})

	if err != nil {
		return err
	}

	// Create Wallet
	err = walletRepo.Create(userID)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *AuthService) Login(
	hp string,
	pin string,
) (*models.User, error) {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	userRepo := models.NewUserRepository(tx)

	user, err := userRepo.FindByPhone(hp)

	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if user.StatusAccount != "active" {
		return nil, errors.New("akun tidak aktif")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Pin),
		[]byte(pin),
	)

	if err != nil {
		return nil, errors.New("PIN salah")
	}

	return user, nil
}