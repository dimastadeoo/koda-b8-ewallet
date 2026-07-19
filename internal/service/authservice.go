package service

import (
	"context"
	"errors"
	// "time"

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
	sessionRepo := models.NewSessionRepository(tx)
	userLogRepo := models.NewUserLogRepository(tx)

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

	// add session
	session, err := sessionRepo.Create("register")
	err = userLogRepo.Create(
		userID,
		session.ID,
		"Register akun",
		"127.0.0.1",
	)
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
) (*models.LoginSession, error) {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	userRepo := models.NewUserRepository(tx)
	profileRepo := models.NewProfileRepository(tx)
	walletRepo := models.NewWalletRepository(tx)
	sessionRepo := models.NewSessionRepository(tx)
	userLogRepo := models.NewUserLogRepository(tx)

	// Cari user
	user, err := userRepo.FindByPhone(hp)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// Validasi PIN
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Pin),
		[]byte(pin),
	)

	if err != nil {
		return nil, errors.New("PIN salah")
	}

	// Ambil Profile
	profile, err := profileRepo.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	// Ambil Wallet
	wallet, err := walletRepo.FindByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	// Simpan Session ke database
	session, err := sessionRepo.Create("login")
	if err != nil {
		return nil, err
	}

	// Simpan User Log
	err = userLogRepo.Create(
		user.ID,
		session.ID,
		"Login berhasil",
		"127.0.0.1",
	)

	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return &models.LoginSession{
		Session: *session,
		User:    *user,
		Profile: *profile,
		Wallet:  *wallet,
	}, nil
}

func (s *AuthService) ForgotPin(
	email string,
	hp string,
	newPin string,
	// dateBirth *time.Time,
) error {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	userRepo := models.NewUserRepository(tx)
	sessionRepo := models.NewSessionRepository(tx)
	userLogRepo := models.NewUserLogRepository(tx)
	// profileRepo := models.NewProfileRepository(tx)

	// Cari user
	user, err := userRepo.FindByEmailAndPhone(
		email,
		hp,
	)

	
	if err != nil {
		return errors.New("email atau nomor HP tidak ditemukan")
	}
	
	if user.StatusAccount != "active" {
		return errors.New("akun tidak aktif")
	}

	// profile, err := profileRepo.FindByUserID(user.ID)

	// Hash PIN baru
	hashPin, err := bcrypt.GenerateFromPassword(
		[]byte(newPin),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	// Update PIN
	err = userRepo.UpdatePin(
		user.ID,
		string(hashPin),
	)

	if err != nil {
		return err
	}

	// Session
	session, err := sessionRepo.Create("forgot-pin")
	if err != nil {
		return err
	}

	// User Log
	err = userLogRepo.Create(
		user.ID,
		session.ID,
		"Reset PIN",
		"127.0.0.1",
	)

	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
