package main

import "fmt"

func main() {
	var mySlice []int = []int{1, 2, 3}
	for _, value := range mySlice {
		fmt.Println(value)
	}
}
