package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)
	channel := make(chan string, 10) // Use a buffered channel for one of the goroutines
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			ch <- "Hello"
			time.Sleep(2 * time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			channel <- "Rahul"
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		wg.Wait()
		close(ch)
		close(channel)
	}()
	for a := range ch {
		fmt.Println(a)
	}
	for b := range channel {
		fmt.Println(b)
	}
}
