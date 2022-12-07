package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  43,
	}

	wTmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(wTmpFile.Name())

	err = json.NewEncoder(wTmpFile).Encode(person) //using io.Reader-io.Writer interface
	if err != nil {
		panic(err)
	}

	err = wTmpFile.Close()
	if err != nil {
		panic(err)
	}

	rTmpFile, err := os.Open(wTmpFile.Name())
	if err != nil {
		panic(err)
	}

	var rPerson Person
	err = json.NewDecoder(rTmpFile).Decode(&rPerson)
	if err != nil {
		panic(err)
	}

	err = rTmpFile.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(rPerson)

}
