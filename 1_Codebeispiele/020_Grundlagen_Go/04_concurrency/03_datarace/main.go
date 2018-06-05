package main

import (
	"fmt"
	"sync"
)

func main() {
	datarace(6)
}

func datarace(nr int) {
	var a int
	var wg sync.WaitGroup
	wg.Add(nr)
	for i := 0; i < nr; i++ {
		go func(nri int) {
			a = nri
			wg.Done()
		}(i)
	}
	fmt.Println(a)
}
