package main

import (
	"fmt"

	"github.com/0Zyablik0/storage"
)

func main() {

	var s storage.Storage
	fmt.Println("s == nil: ", s == nil)
	fmt.Printf("type of s: %T\n", s)

	s = storage.NewMemoryStorage()
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)

	s = storage.NewDumbStorage()
	fmt.Println("s", s)
	fmt.Printf("type of s: %T\n", s)

	s = nil
	fmt.Println("s: ", s)
	fmt.Printf("type of s: %T\n", s)

	// ----------------------------------------------------------------

	ms := storage.NewMemoryStorage()
	ds := storage.NewDumbStorage()

	spawnEmployees(ms, 5)
	fmt.Println(ms.Get(3))

	spawnEmployees(ds, 3)

}

func spawnEmployees(s storage.Storage, n int) {
	for i := 0; i < n; i++ {
		s.Insert(storage.Employee{Id: i})
	}

}
