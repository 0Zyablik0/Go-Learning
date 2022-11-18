package main

import (
	"errors"
	"fmt"
)

type employee struct {
	id     int
	names  string
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (e employee, err error)
	delete(id int) error
}

type memoryStorage struct {
	employees map[int]employee
}

type dumbStorage struct{}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{employees: make(map[int]employee)}
}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *memoryStorage) insert(e employee) error {
	s.employees[e.id] = e
	return nil
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("Insert of employee with id: %d is complete\n", e.id)
	return nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.employees, id)
	return nil

}
func (s *dumbStorage) delete(id int) error {
	fmt.Printf("Delete of employee with id: %d is complete\n", id)
	return nil

}

func (s *memoryStorage) get(id int) (e employee, err error) {
	e, exists := s.employees[id]
	if !exists {
		return employee{}, errors.New("employee not found")
	}

	return e, nil

}

func (s *dumbStorage) get(id int) (e employee, err error) {
	e = employee{id: id}

	return e, nil
}

func main() {

	var s storage
	fmt.Println("s == nil: ", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = newMemoryStorage()
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)

	s = newDumbStorage()
	fmt.Println("s", s)
	fmt.Printf("type of s: %T\n", s)

	s = nil
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)

	// ----------------------------------------------------------------

	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms, 5)
	fmt.Println(ms.get(3))

	spawnEmployees(ds, 3)

}

func spawnEmployees(s storage, n int) {
	for i := 0; i < n; i++ {
		s.insert(employee{id: i})
	}

}
