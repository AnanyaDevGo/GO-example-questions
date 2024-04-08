package main

import "fmt"

func main() {
	type person struct {
		id    int
		name  string
		email string
	}
	var p1 = person{1, "Rahul", "rahulchacko7@gmail.com"}
	var p2 = person{2, "Ashik", "ashik1@gmail.com"}
	var p3 = person{2, "Ananya", "ananya@gmail.com"}
	fmt.Println("struct1")
	fmt.Println(p1.id)
	fmt.Println(p1.name)
	fmt.Println(p1.email)
	fmt.Println("struct2")
	fmt.Println(p2.id)
	fmt.Println(p2.name)
	fmt.Println(p2.email)
	fmt.Println("struct3")
	fmt.Println(p3.id)
	fmt.Println(p3.name)
	fmt.Println(p3.email)

}
