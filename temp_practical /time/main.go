package main

import (
	"fmt"
	"time"
)

func main() {
	currenttime := time.Now()
	func() {
		for {
			fmt.Println(currenttime.Format("2006-01-02 15:04:05"))
			time.Sleep(1 * time.Second)
		}
	}()
}
