// Variable/ Int/Float in Go

package main

import "fmt"

func main() {
	var nameOne string = "Garv"
	var nameTwo = "Tambi"
	var nameThree string
	var nameFour = 23
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	nameOne = "Garv"
	nameThree = "Tambi"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	nameFive := "GarvTambi"
	fmt.Println(nameFive)

	// bits and memory
	var nameSix int8 = 05
	var nameSeven int8 = -15
	var nameEight uint8 = 25
	nameNine := 30
	fmt.Println(nameSix, nameSeven, nameEight, nameNine)
}
