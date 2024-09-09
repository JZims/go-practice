package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 12
	numbers[1] = 164
	numbers[2] = 93
	numbers[3] = 41
	numbers[4] = 78
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
