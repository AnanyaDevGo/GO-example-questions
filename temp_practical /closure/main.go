package main

import "fmt"

func main() {
	inner := outer()
	fmt.Println(inner())
	fmt.Println(inner())
	fmt.Println(inner())
	inner := outer()
	fmt.Println(inner())
}

func outer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
