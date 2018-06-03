package main

import "fmt"

type name struct {
	Vorname  string
	Nachname string
}

type adresse struct {
	Strasse string
	Ort     string
}

type user struct {
	name
	adresse
}

func main() {
	max := user{
		name{"Max", "Muster"},
		adresse{"Musterstrasse", "Bonn"},
	}
	fmt.Println(max.Vorname, max.name.Vorname)
}
