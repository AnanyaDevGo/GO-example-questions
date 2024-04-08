package main

import "fmt"

func main() {
	var String1, String2, a, b string
	fmt.Println("Enter 2 strings.....")
	fmt.Println("Enter the first string :")
	fmt.Scanf("%s", &String1)
	fmt.Println("Enter the second string :")
	fmt.Scanf("%s", &String2)
	a, b = swap(String1, String2)
	fmt.Println(a, b)
}
func swap(x, y string) (string, string) {
	return y, x
}
