package main

import "fmt"

type Employee struct {
	Name     string
	Position string
	Salary   float64
}

func UpdateSalary(emp *Employee, UpdatedSalary float64) {
	emp.Salary = UpdatedSalary
}

func main() {
	emp := Employee{
		Name:     "Arzoo Joshi",
		Position: "Architect",
		Salary:   50000,
	}
	fmt.Println("Before anunal review", emp)
	UpdateSalary(&emp, 60000)
	fmt.Println("After update", emp)
	fmt.Println("emt.salary", &emp.Salary)
}
