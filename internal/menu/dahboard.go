package menu

import (
	"fmt"

	"github.com/dimastadeoo/koda-b8-ewallet/internal/models"
	"github.com/dimastadeoo/koda-b8-ewallet/internal/utils"
)

func Dashboard(user *models.User) {

	for {

		utils.CallClear()

		fmt.Println("===================")
		fmt.Println("WELCOME")
		fmt.Println(user.HPNumber)
		fmt.Println("===================")

		fmt.Println("1. Show Balance")
		fmt.Println("2. List User")
		fmt.Println("3. Transfer")
		fmt.Println("4. History")
		fmt.Println("5. Logout")

		menu := utils.Input("Choose : ")

		switch menu {

		case "5":
			return
		}

	}

}