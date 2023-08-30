// write a program to add two numbers in go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (int, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)
	num, err := strconv.Atoi(input)
	return num, err
}

func add_number(n int, m int) int {
	sum := n + m
	return sum
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add two numbers in go")
	first_number, err1 := getInput("Enter First Number: ", reader)
	second_number, err2 := getInput("Enter Second Number: ", reader)
	if err1 != nil {
		fmt.Println("The Price Must be a number", err1)
		return
	}
	if err2 != nil {
		fmt.Println("The Price Must be a number")
		return
	}

	sum := add_number(first_number, second_number)

	fmt.Printf("Addition of two number is: %v\n", sum)
}
