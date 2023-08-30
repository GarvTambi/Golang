// Pass by Value  // Pointers

package main

import "fmt"

// Group A -> Strings, Ints, Floats, Booleans, Arrays, Structs -> Non Pointer values
// Group B -> Slices, Maps, Functions -> Pointer Wrapper Values

func updateName(n string) string {
	// n = "Tambi"
	na := n
	na = "Tambi"
	return na
}

func updateMenu(y map[string]float64) {
	y["coffee"] = 2.99
}

func main() {
	name := "Garv"
	name = updateName(name)

	fmt.Println(name)
	fmt.Println(&name) // by Address
	ne := &name
	fmt.Println(*ne) // by value

	menu := map[string]float64{
		"pie":       5.95,
		"ice_cream": 6.10,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
