// User Input

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the billl name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println(name)
	b := newBill(name)
	fmt.Println("Created BIll -", b.name)

	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	mybill := createBill()
	fmt.Println(mybill)
	name, _ := getInput("Item Name: ", reader)
	price, _ := getInput(("Item price: "), reader)
	_, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println("The Price Must be a number")
	}
	fmt.Println(name)
}
