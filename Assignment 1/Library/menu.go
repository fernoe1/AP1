package Library

import (
	"fmt"

	"github.com/fernoe1/Assignment1_TemirlanSapargali/Library/model"
)

func Start() {
	library := model.Library{}
	flag := true

	for flag {
		fmt.Println("[1] Add Book")
		fmt.Println("[2] Borrow Book")
		fmt.Println("[3] Return Book")
		fmt.Println("[4] List Books")
		fmt.Println("[5] Return")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter title")
			var title string
			fmt.Scanln(&title)
			fmt.Println("Enter author")
			var author string
			fmt.Scanln(&author)
			book := model.CreateBook(title, author)

			err := library.AddBook(&book)
			if err != nil {
				fmt.Println(err)
			}

		case 2:
			fmt.Println("Enter title")
			var title string
			fmt.Scanln(&title)
			book, err := library.BorrowBook(title)

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(book)
			}

		case 3:
			fmt.Println("Enter title")
			var title string
			fmt.Scanln(&title)
			err := library.ReturnBook(title)

			if err != nil {
				fmt.Println(err)
			}

		case 4:
			books, err := library.ListAvailableBooks()

			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(books)
			}

		case 5:
			flag = false

		default:
			fmt.Println("Invalid choice")
		}
	}
}
