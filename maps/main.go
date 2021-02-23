package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
	}

	numbers := make(map[int]string)
	numbers[1] = "one"
	numbers[2] = "two"

	delete(colors, "red")
	printMap(numbers)
	fmt.Println(colors)

}

func printMap(m map[int]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
