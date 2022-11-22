package main

import "fmt"

func main() {

	var nilMap map[string]int
	fmt.Println("nilMap is nil: ", nilMap == nil)

	testMap := map[string]int{}
	fmt.Println("tetMap is nil: ", testMap == nil)

	teams := map[string][]string{
		"Orcas":   []string{"Fred", "Ralph", "Bijou"},
		"Lions":   []string{"Sarah", "Peter", "John"},
		"Kittens": []string{"Waldo", "Daniel", "Margo"},
	}

	fmt.Println(teams)

	mapMakeTest := make(map[int]string, 10)
	fmt.Println(mapMakeTest, len(mapMakeTest))

	// comma ok idiom

	m := map[string]int{
		"hello": 5,
		"world": 1,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["!"]
	fmt.Println(v, ok)

	// Deleting from map

	delete(m, "hello")
	fmt.Println(m)

}
