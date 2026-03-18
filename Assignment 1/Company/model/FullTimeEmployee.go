package model

import "strconv"

type FullTimeEmployee struct {
	Name, Surname string
	Salary        float64
}

func (e FullTimeEmployee) GetDetails() string {
	return e.Name + " " + e.Surname + " Salary: " + strconv.FormatFloat(e.Salary, 'f', -1, 64)
}
