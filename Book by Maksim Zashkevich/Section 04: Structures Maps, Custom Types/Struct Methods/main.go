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
	employee1 := newEmployee("John", 20, "Manager", 1000)
	employee2 := newEmployee("Michael", 40, "Worker", 3000)

	employee1.printInfo()
	employee2.printInfo()

	employee1.setName("Victor")
	employee1.printInfo()

}

func newEmployee(name string, age int, position string, salary int) employee {
	return employee{
		name:     name,
		salary:   salary,
		position: position,
		age:      age,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("name: %s, age: %d, position: %s, salary: %d", e.name, e.age, e.position, e.salary)

}

func (e employee) printInfo() {
	fmt.Println(e.getInfo())
}

func (e *employee) setName(name string) {
	e.name = name
}
