// Booleans and Conditions

package main

import "fmt"

func main() {
	age := 45
	fmt.Println(age <= 50)
	fmt.Println(age == 50)
	fmt.Println(age >= 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("The age is less than 30")
	} else if age < 40 {
		fmt.Println("The age is less than 40")
	} else {
		fmt.Println("The age is more than or equal to 40")
	}

	names := []string{"Garv", "Tambi", "Anshul", "Samyak", "Samarth"} // slice array because we didn't specify the size of array

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("The value of index %v is %v\n", index, value)
	}

}
