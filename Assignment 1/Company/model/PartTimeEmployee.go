package model

import "strconv"

type PartTimeEmployee struct {
	Name, Surname string
	Salary        float64
}

func (e PartTimeEmployee) GetDetails() string {
	return e.Name + " " + e.Surname + " Salary: " + strconv.FormatFloat(e.Salary, 'f', -1, 64)
}
