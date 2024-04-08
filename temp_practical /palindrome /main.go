package main

import "fmt"

// func main() {
// 	var num, reverse, remainder, original int

// 	// Read the number from the user
// 	fmt.Print("Enter a number: ")
// 	fmt.Scan(&num)

// 	original = num // Store the original number

// 	// Reverse the digits of the number
// 	for {
// 		remainder = num % 10
// 		reverse = reverse*10 + remainder
// 		num /= 10

// 		if num == 0 {
// 			break
// 		}
// 	}

//		// Check if it's a palindrome
//		if original == reverse {
//			fmt.Printf("%d is a palindrome.\n", original)
//		} else {
//			fmt.Printf("%d is not a palindrome.\n", original)
//		}
//	}
func main() {
	var name string

	fmt.Println("Enter a String")
	fmt.Scan(&name)
	left := 0
	right := len(name) - 1
	for left < right {
		if name[left] != name[right] {
			fmt.Println("String is not palindrome")
			return
		}
		left++
		right--
	}
	fmt.Println("Sting is palindrome")
}
