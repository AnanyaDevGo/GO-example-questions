package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start of the program")

	var wg sync.WaitGroup

	// Launch two goroutines
	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine 1")
	}()
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Goroutine 2")
	}()

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("After WaitGroup")

	// WaitGroup is used to wait for multiple goroutines to complete before proceeding.
}
