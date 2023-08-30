// Struct type

package main

import "fmt"

// Define a struct type named "Person"
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Creating a struct instance
	person1 := Person{
		FirstName: "Garv",
		LastName:  "Tambi",
		Age:       15,
	}

	// Accessing struct fields
	fmt.Println("First Name:", person1.FirstName)
	fmt.Println("Last Name:", person1.LastName)
	fmt.Println("Age:", person1.Age)

	// Modifying struct fields
	person1.Age = 23
	fmt.Println("Updated Age:", person1.Age)
	fmt.Println(person1)
}
