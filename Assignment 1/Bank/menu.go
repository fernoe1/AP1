package Bank

import (
	"fmt"

	"github.com/fernoe1/Assignment1_TemirlanSapargali/Bank/Model"
)

func Start() {
	bank := Model.BankAccount{
		Balance:      0.0,
		Transactions: make([]string, 0),
	}

	flag := true
	for flag {
		fmt.Println("[1] Deposit")
		fmt.Println("[2] Withdraw")
		fmt.Println("[3] Balance || Transactions")
		fmt.Println("[4] Exit")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter amount:")
			var amount float64
			fmt.Scanln(&amount)

			err := bank.Deposit(amount)
			if err != nil {
				fmt.Println(err)
			}

		case 2:
			fmt.Println("Enter amount:")
			var amount float64
			fmt.Scanln(&amount)

			_, err := bank.Withdraw(amount)
			if err != nil {
				fmt.Println(err)
			}

		case 3:
			bal, transactions := bank.GetBalance()
			fmt.Println("Balance:", bal)
			fmt.Println("Transactions:", transactions)

		case 4:
			flag = false

		default:
			fmt.Println("Invalid choice")
		}
	}
}
