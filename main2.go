/// Printing and Formatting Strings

package main

import "fmt"

func main() {
	age := 23
	name := "Garv Tambi"

	fmt.Println("My name is", name, "and age is", age)

	// Printf Formatted %_ -> format specifier
	fmt.Printf("My name is %v and age is %v \n", name, age)
	fmt.Printf("Age is type of %T \n", age)

	// Sprintf Formatted
	var str = fmt.Sprintf("My name is %v and age is %v \n", name, age)
	fmt.Println(str)
}
