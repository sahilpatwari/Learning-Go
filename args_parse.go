//Simple Program to get used to CLI Argument Parsing in Go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := os.Args

	if len(s) == 2 && s[1] == "help" {
		fmt.Println("List of Commands:")
		fmt.Println("add : To add two numbers")
		fmt.Println("sub : To subtract two numbers")
		fmt.Println("multiply : To multiply two numbers")
		fmt.Println("divide : To divide two numbers")
		fmt.Println("Command Syntax :")
		fmt.Println("[Executable] [Command] [number1] [number2]")
		fmt.Println("Example : ./a.exe add 2 3")
		return
	}

	if len(s) != 4 {
		fmt.Printf("Undefined Command! Please use help to get the list of commands")
		return
	}
    

	num1, err1 := strconv.Atoi(s[2])
	num2, err2 := strconv.Atoi(s[3])
	if err1 != nil {
		fmt.Println("Error: ",err1)
		return
	}

	if err2 != nil {
		fmt.Println("Error: ",err2)
		return
	}
    
	switch s[1] {
	case "add":
		fmt.Printf("Sum : %v", num1 + num2)
	case "sub":
		fmt.Printf("Difference : %v", num1 - num2)
	case "multiply":
		fmt.Printf("Product : %v", num1 * num2)
	case "divide":
		if num2 != 0 {
		   fmt.Printf("Integer Quotient : %v", num1 / num2)
		} else {
			fmt.Println("Error! Division by Zero")
		}
	default:
		fmt.Printf("Undefined Command! Please use help to get the list of commands")
	}
}