package main

import "fmt"

func main() {
	var f flächer
	f = dreieck{2, 3}
	fmt.Println(fläche(f))
	d := dreieck{4, 5}
	fmt.Println(fläche(d), d)
}

type flächer interface {
	fläche() int
}

func fläche(f flächer) int {
	return f.fläche()
}

type dreieck struct {
	b, h int
}

func (d dreieck) fläche() int {
	return d.b * d.h / 2
}

func (d dreieck) String() string {
	return fmt.Sprintf("Dreieck: b:%d, h:%d", d.b, d.h)
}
