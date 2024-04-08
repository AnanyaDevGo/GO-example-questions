package main

import "fmt"

func main() {
	dist := make(map[string]int)
	dist["Kottayam"] = 10
	dist["Trissur"] = 20
	dist["Ernakulam"] = 30
	fmt.Println(dist)
	fmt.Println(dist["Kottayam"])
	key := "Kollam"
	if val, ok := dist[key]; ok {
		dist[key] = val + 10
		fmt.Println(dist[key])
	} else {
		fmt.Println("Key is not present")
	}
}
