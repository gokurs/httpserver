package main

import "fmt"

func main() {
	r := rechteck{5, 10}
	d := dreieck{5, 10}
	fmt.Println("r:", flaeche(r))
	fmt.Println("d:", flaeche(d))
}

type form interface {
	flaeche() int
}

func flaeche(f form) int {
	return f.flaeche()
}

type rechteck struct {
	l, b int
}

func (r rechteck) flaeche() int {
	return r.l * r.b
}

type dreieck struct {
	l, b int
}

func (d dreieck) flaeche() int {
	return d.l * d.b / 2
}
