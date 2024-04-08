package main

import "fmt"

func sumArray(arr []int, result chan<- int) {
	sum := 0

	for _, num := range arr {
		sum += num
	}
	result <- sum
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := make(chan int, 2)

	mid := len(arr) / 2
	fHalf := arr[mid:]
	sHalf := arr[:mid]
	go sumArray(fHalf, result)
	go sumArray(sHalf, result)

	sum1 := <-result
	sum2 := <-result

	totalSum := sum1 + sum2

	fmt.Println("Total : ", totalSum)

}
