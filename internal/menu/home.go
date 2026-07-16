package menu

import (
	"fmt"
	"strings"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/service"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/utils"
)

type HomeMenu struct {
	authService *service.AuthService
	userService *service.UserService
	trasferService *service.TransferService
	walletService *service.WalletService

}

func NewHomeMenu(
	authService service.AuthService,
	userService service.UserService,
	trasferService service.TransferService,
	walletService service.WalletService,

	) *HomeMenu {
	return &HomeMenu{
		authService: &authService,
		userService: &userService,
		trasferService: &trasferService,
		walletService: &walletService,
	}
}

func (h *HomeMenu) Home() {

	for {

		utils.CallClear()

		fmt.Println("============================")
		fmt.Println("        E-WALLET")
		fmt.Println("============================")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("0. Exit")
		fmt.Println("============================")

		menu := utils.Input("Choose : ")

		switch menu {

		case "1":
			h.login()

		case "2":
			h.register()

		case "0":
			fmt.Println("Terima kasih telah menggunakan E-Wallet")
			return

		default:
			fmt.Println("Menu tidak tersedia")
			utils.PressEnter("Tekan Enter...")
		}
	}
}

func (h *HomeMenu) register() {

	utils.CallClear()

	fmt.Println("========== REGISTER ==========")

	name := strings.TrimSpace(utils.Input("Nama           : "))
	email := strings.TrimSpace(utils.Input("Email          : "))
	hp := strings.TrimSpace(utils.Input("No HP          : "))
	gender := strings.ToUpper(strings.TrimSpace(utils.Input("Gender (M/F)   : ")))
	pin := strings.TrimSpace(utils.Input("PIN (6 digit)  : "))

	if name == "" || email == "" || hp == "" || gender == "" || pin == "" {
		fmt.Println("\nSemua field wajib diisi.")
		utils.PressEnter("Tekan Enter...")
		return
	}

	if gender != "M" && gender != "F" {
		fmt.Println("\nGender hanya boleh M atau F.")
		utils.PressEnter("Tekan Enter...")
		return
	}

	if len(pin) != 6 {
		fmt.Println("\nPIN harus 6 digit.")
		utils.PressEnter("Tekan Enter...")
		return
	}

	err := h.authService.Register(
		email,
		hp,
		pin,
		name,
		gender,
	)

	if err != nil {
		fmt.Println("\nRegister gagal :", err)
	} else {
		fmt.Println("\nRegister berhasil.")
	}

	utils.PressEnter("Tekan Enter...")
}

func (h *HomeMenu) login() {

	utils.CallClear()

	fmt.Println("=========== LOGIN ===========")

	hp := strings.TrimSpace(utils.Input("No HP : "))
	pin := strings.TrimSpace(utils.Input("PIN   : "))

	if hp == "" || pin == "" {
		fmt.Println("\nNo HP dan PIN wajib diisi.")
		utils.PressEnter("Tekan Enter...")
		return
	}

	user, err := h.authService.Login(hp, pin)

	if err != nil {
		fmt.Println("\nLogin gagal :", err)
		utils.PressEnter("Tekan Enter...")
		return
	}

	h.dashboard(user)
}

func (h *HomeMenu) dashboard(user *models.LoginSession) {
	h.Dashboard(user)
}


