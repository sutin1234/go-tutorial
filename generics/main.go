package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7}
	sum := sum(numbers)
	fmt.Println(sum)
}

type SUMIntOrFLOAT interface {
	int | float64
}

func sum[G SUMIntOrFLOAT](obj []G) G {
	var result G
	for _, item := range obj {
		result += item
	}
	return result
}
