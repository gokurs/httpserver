package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	go do(c, 1)
	go do(c, 3)
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
