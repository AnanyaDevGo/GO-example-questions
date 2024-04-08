package main

import "fmt"

func main() {
	var number [10]int
	number[0] = 100
	number[1] = 101
	number[2] = 102
	number[3] = 103
	number[4] = 104
	fmt.Println(number)
	fmt.Println(number[2], number[5])
	fmt.Println("Length of the array is : ", len(number))
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}
}
