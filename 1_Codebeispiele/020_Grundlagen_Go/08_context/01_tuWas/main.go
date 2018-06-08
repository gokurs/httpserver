package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()
	go tuWas(ctx)
	wg.Wait()

}

func tuWas(ctx context.Context) {
	defer wg.Done()
	done := make(chan bool)
	go func() {
		fmt.Println("Fange an")
		anstrengendeOperation()
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Fertig mit anstrengenderOperation")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}

func anstrengendeOperation() {
	i := rand.Intn(100)
	fmt.Println("i:", i)
	time.Sleep(time.Millisecond * time.Duration(i))
}
