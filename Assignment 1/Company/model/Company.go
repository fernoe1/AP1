package model

type Company struct {
	employees map[int]Employee
}

var IncrementingID int = 1

func (c *Company) initialize() {
	c.employees = make(map[int]Employee)
}

func (c *Company) AddEmployee(emp Employee) {
	if c.employees == nil {
		c.initialize()
	}

	c.employees[IncrementingID] = emp
	IncrementingID++
}

func (c *Company) ListEmployees() []Employee {
	employeesSlice := make([]Employee, 0)
	for _, employee := range c.employees {
		employeesSlice = append(employeesSlice, employee)
	}

	return employeesSlice
}
