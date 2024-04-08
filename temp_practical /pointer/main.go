package main

import "fmt"

func main() {
	var number int = 100
	fmt.Println(number)
	var ptr *int
	ptr = &number
	fmt.Println(ptr)
	*ptr = 200
	fmt.Println(number)
}
