package main

import "fmt"

type user struct {
	Vorname  string
	Nachname string
	Wohort   string
	Alter    int
}

func main() {
	max := user{"Max", "Muster", "Bonn", 23}
	moritz := user{
		Nachname: "Mustermann",
		Vorname:  "Moritz",
	}
	fmt.Println(max)
	fmt.Println(moritz)
}
