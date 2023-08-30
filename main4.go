// The Standard Library

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	message := "Hello There friends!"

	fmt.Println(strings.Contains(message, "Hello"))
	fmt.Println(strings.Contains(message, "hello"))
	fmt.Println(strings.ReplaceAll(message, "Hello", "Howdy"))
	fmt.Println(strings.ToUpper(message))
	fmt.Println(strings.Index(message, "Th"))
	fmt.Println(strings.Split(message, " "))

	ages := []int{2, 5, 3, 6, 5, 4, 7, 1}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 4)
	fmt.Println(index)

	names := []string{"Garv", "Tambi"}
	// names := [2]string{"Garv", "Tambi"}    // This will give error
	fmt.Println(sort.SearchStrings(names, "Tambi"))
}
