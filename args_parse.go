//Simple Program to get used to CLI Argument Parsing in Go
package main

import (
	"fmt"
	"flag"
)

func main() {
    
	var opFlag = flag.String("op","undefined","Operations allowed : add, sub, multiply, divide")
	var numFlag1 = flag.Int("num1",0,"First Number")
	var numFlag2 = flag.Int("num2",0,"Second Number")
    
	flag.Parse()
	op , num1 , num2 := *opFlag , *numFlag1 , *numFlag2
	switch op {
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