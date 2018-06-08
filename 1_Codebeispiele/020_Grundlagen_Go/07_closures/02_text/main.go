package main

import (
	"fmt"
)

type myFunc func(int) string

func intToString(i int) string {
	return fmt.Sprintf("%d", i)
}

func addPrefix(s string, f myFunc) myFunc {
	return func(i int) string {
		return fmt.Sprintf("%s%s", s, f(i))
	}
}
func main() {
	f := addPrefix("--->", myFunc(intToString))
	fmt.Println(f(1))
}
