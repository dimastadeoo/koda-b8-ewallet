package menu

import (
	"fmt"
	"strconv"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/utils"
)


func (h *HomeMenu) listUser() {

	users, err := h.userService.GetAll()

	if err != nil {
		fmt.Println(err)
		utils.PressEnter("Enter...")
		return
	}

	fmt.Println("==============================================")
	fmt.Printf("%-3s %-20s %-15s %-25s\n",
		"No", "Nama", "No HP", "Email")
	fmt.Println("==============================================")

	for i, user := range users {

		fmt.Printf("%-3d %-20s %-15s %-25s\n",
			i+1,
			user.Name,
			user.Phone,
			user.Email,
		)

	}

	utils.PressEnter("\nEnter...")
}

func (h *HomeMenu) historyTransfer(walletId int64) {

	listTransfer, err := h.trasferService.History(walletId)

	if err != nil {
		fmt.Println(err)
		utils.PressEnter("Enter...")
		return
	}

	fmt.Println("==============================================================================================================")
	fmt.Printf("%-3s | %-20s | %-15s | %-15s | %-10s | %-12s | %-20s\n",
		"No", "Tanggal", "Nama", "No HP", "Status", "Amount", "Notes")
	fmt.Println("==============================================================================================================")

	for i, list := range listTransfer {
		fmt.Printf("%-3d | %-20s | %-15s | %-15s | %-10s | %-12.2f | %-20s\n",
			i+1,
			list.CreatedAt,
			list.Name,
			list.Phone,
			list.Status,
			list.Amount,
			list.Notes,
		)
	}

	fmt.Println("==============================================================================================================")

	utils.PressEnter("\nEnter...")
}

func (h *HomeMenu) transfer(
	session *models.LoginSession,
) {

	phone := utils.Input("No HP Tujuan : ")

	amountStr := utils.Input("Nominal      : ")

	note := utils.Input("Catatan      : ")

	amount, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		fmt.Println("Nominal tidak valid")
		utils.PressEnter("Enter...")
		return
	}

	err = h.trasferService.Transfer(
		session,
		phone,
		amount,
		note,
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Transfer berhasil")
	}

	utils.PressEnter("Enter...")
}

func (h *HomeMenu) topUp(
	session *models.LoginSession,
) {

	utils.CallClear()

	fmt.Println("========== TOP UP ==========")

	input := utils.Input("Nominal : ")

	amount, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Nominal tidak valid")
		utils.PressEnter("Enter...")
		return
	}

	err = h.walletService.TopUp(
		session,
		amount,
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Top Up berhasil")
		fmt.Printf("Saldo sekarang : %.2f\n", session.Wallet.Balance)
	}

	utils.PressEnter("Enter...")
}


func (h *HomeMenu) Dashboard(session *models.LoginSession) {


	for {

		utils.CallClear()

		fmt.Println("==========================")
		fmt.Println("Welcome,", session.Profile.Name)
		fmt.Println("==========================")

		fmt.Println("1. Show Balance")
		fmt.Println("2. List User")
		fmt.Println("3. Transfer")
		fmt.Println("4. History")
		fmt.Println("5. TopUp")
		fmt.Println("0. Keluar")

		menu := utils.Input("Choose : ")

		switch menu {
		case "1":
			fmt.Println("-----------------------------------")
			fmt.Println("Email   :", session.User.Email)
			fmt.Println("HP      :", session.User.HPNumber)
			fmt.Println("Saldo   :", session.Wallet.Balance)
			fmt.Println("-----------------------------------")
			utils.PressEnter("Tekan Enter untuk lanjut...")
		case "2":
			h.listUser()
		case "3":
			h.transfer(session)
		case "4":
			h.historyTransfer(session.Wallet.ID)
		case "5":
			h.topUp(session)
		case "0":
			return
		default:
			fmt.Println("pilihan tidak ada")
			utils.PressEnter("Tekan Enter untuk coba lagi")
		}

	}

}
