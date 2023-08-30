// Fuctions

package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Printf("Hello!, Good Morning %v\n", n)
}
func sayBye(n string) {
	fmt.Printf("Bye %v\n", n)
}
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Garv")
	sayBye("Garv")

	cycleNames([]string{"Garv", "Tambi", "Anshul", "Samyak", "Samarth"}, sayGreeting)
	cycleNames([]string{"Garv", "Tambi", "Anshul", "Samyak", "Samarth"}, sayBye)

	a1 := circleArea(2.5)
	fmt.Println(a1)

	fmt.Printf("circle1 is %0.2f\n", a1)
}
