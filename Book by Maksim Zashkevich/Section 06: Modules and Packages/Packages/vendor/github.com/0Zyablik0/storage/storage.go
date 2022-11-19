package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id     int
	Names  string
	Age    int
	Salary int
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (e Employee, err error)
	Delete(id int) error
}

type memoryStorage struct {
	employees map[int]Employee
}

type dumbStorage struct{}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{employees: make(map[int]Employee)}
}

func NewDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *memoryStorage) Insert(e Employee) error {
	s.employees[e.Id] = e
	return nil
}

func (s *dumbStorage) Insert(e Employee) error {
	fmt.Printf("Insert of employee with id: %d is complete\n", e.Id)
	return nil
}

func (s *memoryStorage) Delete(id int) error {
	delete(s.employees, id)
	return nil

}
func (s *dumbStorage) Delete(id int) error {
	fmt.Printf("Delete of employee with id: %d is complete\n", id)
	return nil

}

func (s *memoryStorage) Get(id int) (e Employee, err error) {
	e, exists := s.employees[id]
	if !exists {
		return Employee{}, errors.New("employee not found")
	}

	return e, nil

}

func (s *dumbStorage) Get(id int) (e Employee, err error) {
	e = Employee{Id: id}

	return e, nil
}
