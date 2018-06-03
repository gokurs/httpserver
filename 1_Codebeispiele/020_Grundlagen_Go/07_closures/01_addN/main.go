package main

import "fmt"

func addN(n int) func(int) int {
	return func(a int) int {
		return a + n
	}
}

func main() {
	add2 := addN(2)
	fmt.Println(add2(4))
	add10 := addN(10)
	fmt.Println(add10(1), add10(10))
	fmt.Println(add2(2))
}
