package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(4 * time.Second)
			fmt.Println("Hello")
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(3 * time.Second)
			fmt.Println("Hai")
		}
	}()
	wg.Wait()
}
