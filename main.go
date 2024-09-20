package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         int      `json:"age,omitempty"`
	Password    string   `json:"-"` // Ignore this field in JSON
	Permissions []string `json:"roles"`
}

func main() {
	u := User{
		ID:          1,
		Name:        "Sathwik",
		Age:         22,
		Password:    "nani", // This field will be ignored in the JSON output
		Permissions: []string{"admin", "grp member"},
	}

	b, err := json.Marshal(u)

	if err != nil {
		fmt.Println("error marshalling JSON", err)
		panic(err)
	}

	fmt.Println(string(b)) // Password field will not be included in the output
}
