// Arrays and Slices

package main

import "fmt"

func main() {
	var ages = [3]int{1, 2, 3}
	fmt.Println(ages, len(ages))

	names := [2]string{"Garv", "Tambi"}
	fmt.Println(names, len(names))

	/// slices
	var scores = []int{5, 10, 12, 20}
	scores[2] = 15
	scores = append(scores, 25)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := ages[1:3]
	rangeTwo := scores[2:]
	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
}
