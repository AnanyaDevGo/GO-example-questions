package main

import "fmt"

func main() {
	str := "abcd dcba"
	runes := []rune(str)
	fmt.Printf("%c", runes)
	runes1 := 65
	fmt.Println(runes1)
}
