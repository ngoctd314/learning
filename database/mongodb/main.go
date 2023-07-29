package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name bool `json:"name,omitempty"`
}

func main() {
	p := person{}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}
