//Simple Program to implement Stack Data Structure in Go

package main

import(
	"errors"
	"fmt"
)
type Stack struct {
	arr []int 
	top int
}

var MAX_CAPACITY = 4096

//Capacity Value should not exceed 4096(Max Limit fot the Stack)
func stack_init(stack *Stack , capacity int) (int , error) {
	if capacity > MAX_CAPACITY {
		return -1 , errors.New("Stack size exceeds Maximum Allowed Capacity!")
	}

	stack.arr = make([]int,capacity)
    stack.top = -1
	return 0 , nil
}

func push(stack *Stack , val int) (int , error) {
	if stack.top == len(stack.arr) - 1 {
		return -1 , errors.New("Stack is Full! Push Operation Not Possible!")
	}
    
	stack.top = stack.top + 1
	stack.arr[stack.top] = val
	return 0 , nil
}

func pop(stack *Stack) (int , error) {
	if stack.top == -1 {
		return -1 , errors.New("Stack is Empty! Pop Operation Not Possible")
	} 
    
	val := stack.arr[stack.top]
	stack.top = stack.top - 1
	return val, nil
}

func top(stack *Stack) (int , error) {
	if stack.top == -1 {
		return -1 , errors.New("Stack is Empty! Top Operation Not Possible")
	}

	return stack.arr[stack.top] , nil
}

func main() {
	fmt.Printf("Enter Capacity for the Stack : ")
	var capacity int
	fmt.Scanf("%d",&capacity)

	var stack *Stack = new(Stack)
	_ , err := stack_init(stack,capacity)

	if err != nil {
		fmt.Print("Error : ",err)
		return
	}
	
	for {
		fmt.Printf("\npush , pop or top??\n")
		var op string
		fmt.Scanf("%s",&op)
		if op == "push" {
			fmt.Printf("Enter value : ")
			var val int
			fmt.Scanf("%d",&val)
			_ , err := push(stack,val)

			if err != nil {
				fmt.Println("Error : ",err)
				break
			}
            
			fmt.Printf("%d pushed onto stack",val)
		} else if op == "pop" {
			ret , err := pop(stack)

			if err != nil {
				fmt.Println("Error : ",err)
				break
			}

			fmt.Printf("%d popped from stack",ret)
		} else if op == "top" {
			ret , err := top(stack)

			if err != nil {
				fmt.Println("Error : ",err)
				break
			}

			fmt.Printf("%d top value of stack",ret)
		}
	}
}