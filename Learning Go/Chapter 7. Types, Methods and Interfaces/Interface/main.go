package main

import "fmt"

type LogicProvider struct{}

func (lp LogicProvider) Process(data string) string {
	return "processed data"
}

type Logic interface {
	Process(data string) string
}

type Client struct {
	L Logic
}

func (c Client) Program() {
	data := "data"
	fmt.Println(data)
	fmt.Println(c.L.Process(data))

}

func main() {
	c := Client{
		L: LogicProvider{},
	}
	c.Program()
}
