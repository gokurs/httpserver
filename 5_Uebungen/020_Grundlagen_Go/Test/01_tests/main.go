/*
	Ändere den Code hier, sodass der Test bestanden wird
*/
package main

import "fmt"

func main() {
	var s string
	s = getHelloString()
	fmt.Println(s)
}

func getHelloString() string {
	return "Hello World!"
}
