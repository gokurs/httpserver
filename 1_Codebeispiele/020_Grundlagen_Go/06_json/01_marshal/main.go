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
	alter    int    `json:"alter,omitempty"`
}

func main() {
	max := User{"Max", "Muster", "Berlin", 66}
	json, err := json.MarshalIndent(max, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ACHTUNG: Warum wird 'alter' nicht exportiert?")
	fmt.Println(string(json))
}
