package main

import (
	"bytes"
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
	buf := &bytes.Buffer{} // implementiert io.Reader und io.Writer
	err := json.NewEncoder(buf).Encode(max)
	if err != nil {
		log.Fatal(err)
	}
	var u User
	dec := json.NewDecoder(buf)
	err = dec.Decode(&u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}
