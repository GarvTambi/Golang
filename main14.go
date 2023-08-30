// write a program to find if string is palindrome

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func check_string(n string) string {
	m := len(n)
	for i := 0; i < m/2; i++ {
		if n[i] == n[m-1-i] {
			continue
		} else {
			return "not Palindrome"
		}
	}
	return "Palindrome"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Check string is palindrome or not")
	enter_string, err := getInput("Enter the string: ", reader)
	if err != nil {
		fmt.Println("Enter valid string", err)
		return
	}

	check_string := check_string(enter_string)
	fmt.Printf("The Input String is %v\n", check_string)
}
