package main

type Foo struct {
	A string `json:"a,omitempty"`
	B string `json:"b,omitempty"`
}

type Bar struct {
	F1 *Foo `json:"f_1,omitempty"`
	F2 Foo  `json:"f_2,omitempty"`
}

func main() {

}
