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

func main() {
	max := struct {
		name
		adresse
	}{
		name{"Max", "Muster"},
		adresse{"Musterstrasse", "Bonn"},
	}
	fmt.Println(max.Vorname, max.name.Vorname)
}
