//Simple Program to print nth Fibonacci Number : Both Recursive and Iterative 
package main 

import "fmt"

// n = 1 means 1st term of Fibonacci Sequence which is 0 and so on
func fibonacciRecursive(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	return fibonacciRecursive(n - 1) + fibonacciRecursive(n - 2)
}

// n = 1 means 1st term of Fibonacci Sequence which is 0 and so on
func fibonacciIterative(n int) int {
	if n <= 1 {
		return 0
	} else if n == 2 {
		return 1
	}

	arr := make([]int,n)
	arr[0], arr[1] = 0 , 1

	for i := 2; i < n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}

	return arr[n - 1]
}
func main() {
	var n int
	fmt.Printf("Enter the number : ")
	fmt.Scanf("%d",&n)
    res := fibonacciIterative(n)
	fmt.Printf("%dth term of the Fibonacci Sequence is (Iterative): %d\n", n, res)
	res = fibonacciRecursive(n)
	fmt.Printf("%dth term of the Fibonacci Sequence is (Recursive): %d\n", n, res)
}