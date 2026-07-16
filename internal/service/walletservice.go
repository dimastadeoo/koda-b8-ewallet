package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/jackc/pgx/v5"
)


type WalletService struct {
	conn *pgx.Conn
}

func NewWalletService(conn *pgx.Conn) *WalletService {
	return &WalletService{
		conn: conn,
	}
}

func (s *WalletService) TopUp(
	login *models.LoginSession,
	amount float64,
) error {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	walletRepo := models.NewWalletRepository(tx)
	transactionRepo := models.NewTransactionRepository(tx)
	walletLogRepo := models.NewWalletLogRepository(tx)

	// Lock wallet
	wallet, err := walletRepo.FindByIDForUpdate(login.Wallet.ID)
	if err != nil {
		return err
	}

	if wallet.Status != "active" {
		return errors.New("wallet tidak aktif")
	}

	if amount <= 0 {
		return errors.New("nominal harus lebih dari 0")
	}

	newBalance := wallet.Balance + amount

	// Update balance
	err = walletRepo.UpdateBalance(
		wallet.ID,
		newBalance,
	)

	if err != nil {
		return err
	}

	// Transaction
	transactionID, err := transactionRepo.Create(models.Transaction{
		WalletID:        wallet.ID,
		TransactionType: "TOPUP",
		Amount:          amount,
		Status:          "SUCCESS",
		Description:     "Top Up Saldo",
	})

	if err != nil {
		return err
	}

	// Wallet Log
	err = walletLogRepo.Create(models.WalletLog{
		WalletID: wallet.ID,
		Action:   "TOPUP",
		Description: fmt.Sprintf(
			"Top Up saldo Rp %.0f (Transaction ID: %d)",
			amount,
			transactionID,
		),
	})

	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	// Update session
	login.Wallet.Balance = newBalance

	return nil
}