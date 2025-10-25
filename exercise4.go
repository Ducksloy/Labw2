package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func multiply(a int, b int) int {
	return a * b
}

func divide(a int, b int) (int, string) {
	if b == 0 {
		return 0, "Error: Division by zero"
	} else {
		return a / b, ""
	}
}

func modulus(a int, b int) (int, string) {
	if b == 0 {
		return 0, "Error: Division by zero"
	} else {
		return a % b, ""
	}
}

func main() {
	var num1, num2 int 
	var operation string
	for {
		fmt.Print("Enter an operation ( + - * / % ) : ")
		fmt.Scan(&operation)
		if operation == "exit" || operation == "ex" || operation == "Exit" || operation == "Ex" {
			break
		}
		fmt.Print("Enter first number: ")
		fmt.Scan(&num1)
		fmt.Print("Enter second number: ")
		fmt.Scan(&num2)

		switch operation {
		case "+":
			fmt.Printf("Result: %d\n", add(num1, num2))
		case "-":
			fmt.Printf("Result: %d\n", subtract(num1, num2))
		case "*":
			fmt.Printf("Result: %d\n", multiply(num1, num2))
		case "/":
			val, err := divide(num1, num2)
			if err == "" {
				fmt.Printf("Result: %d\n", val)
			} else if val == 0 {
				fmt.Printf("%s\n", err)
			}
		case "%":
			val, err := modulus(num1, num2)
			if err == "" {
				fmt.Printf("Result: %d\n", val)
			} else if val == 0 {
				fmt.Printf("%s\n", err)
			}
		default:
			fmt.Println("Invalid operation")
		}
	}
}