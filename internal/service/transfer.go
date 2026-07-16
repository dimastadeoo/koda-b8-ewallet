package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/jackc/pgx/v5"
)

type TransferService struct {
	conn *pgx.Conn
}

func NewTransferService(conn *pgx.Conn) *TransferService {
	return &TransferService{
		conn: conn,
	}
}

func (s *TransferService) Transfer(
	login *models.LoginSession,
	receiverPhone string,
	amount float64,
	note string,
) error {

	ctx := context.Background()

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("conn: %w", err)

	}

	defer tx.Rollback(ctx)

	userRepo := models.NewUserRepository(tx)
	walletRepo := models.NewWalletRepository(tx)
	transactionRepo := models.NewTransactionRepository(tx)
	transferRepo := models.NewTransferRepository(tx)
	walletLogRepo := models.NewWalletLogRepository(tx)

	// ============================
	// Cari User Tujuan
	// ============================

	receiverUser, err := userRepo.FindByPhone(receiverPhone)
	if err != nil {
		return errors.New("user tujuan tidak ditemukan")
	}

	if receiverUser.ID == login.User.ID {
		return errors.New("tidak dapat transfer ke akun sendiri")
	}
	
	// ============================
	// Ambil Wallet
	// ============================

	senderWallet, err := walletRepo.FindByIDForUpdate(login.Wallet.ID)
	if err != nil {
		return fmt.Errorf("FindMyWallet: %w", err)

	}

	receiverWallet, err := walletRepo.FindByUserID(receiverUser.ID)
	if err != nil {
		return fmt.Errorf("FindReceiveWallet: %w", err)
	}


	// ============================
	// Validasi
	// ============================

	if senderWallet.Status != "active" {
		return errors.New("wallet pengirim tidak aktif")
	}

	if receiverWallet.Status != "active" {
		return errors.New("wallet penerima tidak aktif")
	}

	if senderWallet.Balance < amount {
		return errors.New("saldo tidak mencukupi")
	}

	// ============================
	// Hitung Saldo Baru
	// ============================

	newSenderBalance := senderWallet.Balance - amount
	newReceiverBalance := receiverWallet.Balance + amount

	// ============================
	// Update Wallet
	// ============================

	err = walletRepo.UpdateBalance(
		senderWallet.ID,
		newSenderBalance,
	)

	if err != nil {
		return fmt.Errorf("UpdateMyBalance: %w", err)

	}

	err = walletRepo.UpdateBalance(
		receiverWallet.ID,
		newReceiverBalance,
	)

	if err != nil {
		return fmt.Errorf("UpdateReceiveBalance: %w", err)
	}

	// ============================
	// Transaction Sender
	// ============================

	senderTransactionID, err := transactionRepo.Create(models.Transaction{
		WalletID:        senderWallet.ID,
		TransactionType: "TRANSFER OUT",
		Amount:          amount,
		Status:          "SUCCESS",
		Description:     note,
	})

	if err != nil {
		return fmt.Errorf("CreateMyTransac: %w", err)
	}

	// ============================
	// Transaction Receiver
	// ============================

	_, err = transactionRepo.Create(models.Transaction{
		WalletID:        receiverWallet.ID,
		TransactionType: "TRANSFER IN",
		Amount:          amount,
		Status:          "SUCCESS",
		Description:     note,
	})

	if err != nil {
		return fmt.Errorf("CreateReceiveTransac: %w", err)

	}

	// ============================
	// Transfer
	// ============================

	err = transferRepo.Create(models.Transfer{
		TransactionID:    senderTransactionID,
		SenderWalletID:   senderWallet.ID,
		ReceiverWalletID: receiverWallet.ID,
		Notes:            note,
	})

	if err != nil {
		return fmt.Errorf("CreateTransfer: %w", err)

	}

	// ============================
	// Wallet Log Sender
	// ============================

	err = walletLogRepo.Create(models.WalletLog{
		WalletID: senderWallet.ID,
		Action:   "TRANSFER OUT",
		Description: fmt.Sprintf(
			"Transfer %.2f ke %s",
			amount,
			receiverUser.HPNumber,
		),
	})

	if err != nil {
		return fmt.Errorf("CreateMyWalletLog: %w", err)

	}

	// ============================
	// Wallet Log Receiver
	// ============================

	err = walletLogRepo.Create(models.WalletLog{
		WalletID: receiverWallet.ID,
		Action:   "TRANSFER IN",
		Description: fmt.Sprintf(
			"Menerima transfer %.2f dari %s",
			amount,
			login.User.HPNumber,
		),
	})

	if err != nil {
		return fmt.Errorf("CreateReceiveWalletLog: %w", err)

	}

	// ==========================		return err==
	// Commit
	// ============================

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("Commit: %w", err)
	}

	// Update Session
	login.Wallet.Balance = newSenderBalance

	return nil
}


func (s *TransferService) History(
	walletID int64,
) ([]models.TransferHistory, error) {

	tx, err := s.conn.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(context.Background())

	repo := models.NewTransferRepository(tx)

	return repo.History(walletID)
}