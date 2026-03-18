package Company

import (
	"fmt"

	"github.com/fernoe1/Assignment1_TemirlanSapargali/Company/model"
)

func Start() {
	company := model.Company{}
	flag := true

	for flag {
		fmt.Println("[1] Add employee")
		fmt.Println("[2] List employees")
		fmt.Println("[3] Return")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter employee first name")
			var name string
			fmt.Scan(&name)
			fmt.Println("Enter employee last name")
			var lastName string
			fmt.Scan(&lastName)
			fmt.Println("Enter employee salary")
			var salary float64
			fmt.Scan(&salary)
			fmt.Println("Enter employee type")
			fmt.Println("[1] Full time")
			fmt.Println("[2] Part time")
			var employeeType int
			fmt.Scan(&employeeType)

			if employeeType == 1 {
				fullTimeEmployee := model.FullTimeEmployee{
					Name:    name,
					Surname: lastName,
					Salary:  salary,
				}

				company.AddEmployee(fullTimeEmployee)
			} else if employeeType == 2 {
				partTimeEmployee := model.PartTimeEmployee{
					Name:    name,
					Surname: lastName,
					Salary:  salary,
				}

				company.AddEmployee(partTimeEmployee)
			} else {
				fmt.Println("Invalid Employee type")
			}

		case 2:
			employeeSlice := company.ListEmployees()
			fmt.Println(employeeSlice)

		case 3:
			flag = false

		default:
			fmt.Println("Invalid choice")
		}
	}
}
