package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Vorname  string `json:"vorname,omitempty"`
	Nachname string `json:"nachname,omitempty"`
	Wohnort  string `json:"wohnort,omitempty"`
}

func main() {
	max := User{"Max", "Muster", "Berlin"}
	jsonBytes, err := json.MarshalIndent(max, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	var u User
	err = json.Unmarshal(jsonBytes, &u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)

}
