package main

import "fmt"

func main() {
	var num []int
	fmt.Println(num)
	num = append(num, 1, 2, 3, 4, 5)
	fmt.Println(num)
	fmt.Println(len(num))
	fmt.Println(cap(num))
	// for i := 0; i < len(num); i++ {
	// 	fmt.Println("Items on the slice \n", num[i])
	// }
	target := make([]int, len(num))
	copy(target, num) //coping the existing slice to new slice
	fmt.Println(target)
	num[1] = 100
	fmt.Println(num)
	fmt.Println(num[:2])
	fmt.Println(num)
	fmt.Println(num[2:])
}
