package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println(time.Now().Format("15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02"))
	}
	//	t := time.Now()
	//switch t {
	//case t.Hour() < 12:
	//	fmt.Println("Good morning ")
	//	case t.Hour() < 17:
	//	fmt.Println("Good Evening")
	//case t.Hour() < 22:
	//	fmt.Println("Good Night")
	//	default:
	//		fmt.Println("Good Day...!!!")
	//	}
}
