package main

import "fmt"

func main() {
	fmt.Println(naked(17))
}
func naked(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
