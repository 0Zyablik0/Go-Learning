package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	DateOrdered time.Time `json:"date_ordered"`
	CustomerID  string    `json:"customer_id"`
	Items       []Item    `json:"itmes,omitempty"`
}

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	myOrder := Order{
		ID:          "42",
		DateOrdered: time.Now(),
		CustomerID:  "123",
		Items: []Item{
			{
				ID:   "1",
				Name: "Item 1",
			},
			{
				ID:   "2",
				Name: "Item 2",
			},
		},
	}
	data, err := json.Marshal(myOrder)
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println(string(data))
	}

	var unmarshalledOrder Order

	err = json.Unmarshal(data, &unmarshalledOrder)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(unmarshalledOrder)
	}

}
