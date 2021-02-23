package main

import "fmt"

func main() {
	intSlice := createIntSlice(1, 10)
	for _, intValue := range intSlice {
		if intValue%2 == 0 {
			fmt.Println(intValue, "is even")
		} else {
			fmt.Println(intValue, "is odd")
		}
	}
}

func createIntSlice(min int, max int) []int {
	var intSlice []int
	for i := min; i <= max; i++ {
		intSlice = append(intSlice, i)
	}
	return intSlice
}
