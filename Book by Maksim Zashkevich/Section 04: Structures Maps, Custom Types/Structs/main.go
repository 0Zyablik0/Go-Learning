package main

import (
	"fmt"
)

type employee struct {
	name     string
	position string
	age      int
	salary   int
}

func main() {
	basicStructSyntax()

}

func basicStructSyntax() {

	fmt.Println("\n Creating the simple structure: ")

	varName := struct{}{}

	fmt.Printf("varName Type = %T, varName value: %v\n", varName, varName)
	fmt.Println("----------------------------------------------------------------")

	fmt.Println("Creating employee type:")

	employee1 := newEmployee("John", 20, "Manager", 1000)

	employee2 := newEmployee("Michael", 40, "Worker", 3000)

	fmt.Printf("employee1: %+v\n", employee1)
	printEmployee(&employee1)

	fmt.Printf("\nemployee2: %+v\n", employee2)
	printEmployee(&employee2)

	fmt.Println("----------------------------------------------------------------")
}

func newEmployee(name string, age int, position string, salary int) employee {
	return employee{
		name:     name,
		salary:   salary,
		position: position,
		age:      age,
	}
}

func printEmployee(employee *employee) {
	fmt.Printf("Name: %s, age: %d, position: %s, salary: %d\n", employee.name, employee.age, employee.position, employee.salary)

}
