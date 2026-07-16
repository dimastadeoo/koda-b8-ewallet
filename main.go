package main

import (
	"context"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/lib"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/menu"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/service"
)

func main() {

	conn := lib.Conn()
	defer conn.Close(context.Background())

	authService := service.NewAuthService(conn)
	userService := service.NewUserService(conn)
	transferService := service.NewTransferService(conn)
	walletService := service.NewWalletService(conn)

	home := menu.NewHomeMenu(*authService, *userService, *transferService, *walletService)

	home.Home()
}
