package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(total)
}
func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
