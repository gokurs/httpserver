package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go do(c, 10)
	go do(c, 6)
	fmt.Println(<-c)
	fmt.Println("erster Wert erhalten")
	fmt.Println(<-c)
	fmt.Println("zweiter Wert erhalten")

}

func do(c chan<- int, i int) {
	fmt.Println(i, "Wert bekommen")
	time.Sleep(time.Second * time.Duration(i))
	c <- i
}
