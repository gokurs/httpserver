package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go do(c1, 5)
	go do(c2, 2)
	select {
	case i := <-c1:
		fmt.Println(i)
	case i := <-c2:
		fmt.Println(i)
	default:
		fmt.Println("default")
	}
	fmt.Println("Programm ende")
}

func do(c chan int, i int) {
	fmt.Println(i, "gestartet")
	time.Sleep(time.Second * time.Duration(i))
	c <- i
}
