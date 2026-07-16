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

	home := menu.NewHomeMenu(authService)

	home.Home()
}