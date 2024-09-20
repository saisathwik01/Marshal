package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string   `json:"full_name"`
	Age        int      `json:"years_old,omitempty"`
	Occupation string   `json:"-"`
	Languages  []string `json:"spoken_languages"`
}

func main() {

	jsonData := `{"full_name": "Sathwik", "years_old,omitempty":22, "spoken_languages":["Telugu", "English"]}`

	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println("Error unmarshalling JSOn:", err)
		panic(err)
	}

	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Occupation:", person.Occupation)
	fmt.Println("Spoken Languages:", person.Languages)
}
