package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Rahul")
	sillyBoy()
	fmt.Println("Thank You...")
}

func sillyBoy() {
	fmt.Println("Hello Arun")
	SillyGirl()
	fmt.Println("Function Call Ended")
}

func SillyGirl() {
	fmt.Println("Hello Ananya")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered:", r)
		}
	}()
	panic("Ananya Hacked the system..!!!")
	//	fmt.Println("No more execution...")

}

// You can edit this code!
// Click here and start typing.

// import "fmt"

// func main() {
// 	fmt.Println("Simple Example")
// 	check()
// 	fmt.Println("Good Bye...")
// }
// func check() {
// 	fmt.Println("Welcome to Function")
// 	defer func() {
// 		r := recover()
// 		if r != nil {
// 			fmt.Println("Error Recovered...", r)
// 		}
// 	}()
// 	panic("An Eroor encounted...")
// }
