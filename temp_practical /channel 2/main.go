package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(2 * time.Second)
			ch <- "Hello"
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(3 * time.Second)
			ch <- "Hai"
		}
	}()
	timerchan := time.After(10 * time.Second)
	for {
		select {
		case r := <-ch:
			fmt.Println(r)
		case <-timerchan:
			return
		}
	}
}
