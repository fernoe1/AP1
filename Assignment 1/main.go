package main

import (
	"fmt"

	"github.com/fernoe1/Assignment1_TemirlanSapargali/Bank"
	"github.com/fernoe1/Assignment1_TemirlanSapargali/Company"
	"github.com/fernoe1/Assignment1_TemirlanSapargali/Library"
	"github.com/fernoe1/Assignment1_TemirlanSapargali/Shapes"
)

func main() {
	flag := true
	for flag {
		fmt.Println("[1] Library")
		fmt.Println("[2] Shapes")
		fmt.Println("[3] Company")
		fmt.Println("[4] Bank")
		fmt.Println("[0] Exit")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Library.Start()

		case 2:
			Shapes.Start()

		case 3:
			Company.Start()

		case 4:
			Bank.Start()

		case 0:
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}
