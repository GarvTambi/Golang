// Loops

package main

import "fmt"

func main() {
	x := 0

	// while loop
	for x < 5 {
		fmt.Println("The value of x is", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("The value of i is", i)
	}

	names := []string{"Garv", "Tambi", "Anshul", "Samyak", "Samarth"} // slice array because we didn't specify the size of array
	names_length := len(names)
	x = 0

	for x < names_length {
		fmt.Println("value ", names[x])
		x++
	}

	for index, value := range names {
		fmt.Printf("The value of index %v is %v\n", index, value)
		value = "My strings" // This value is not gets upodated in slice= slice is when we don;t specify the size of array
	}

	for _, value := range names {
		fmt.Printf("The value is %v\n", value)
	}
}
