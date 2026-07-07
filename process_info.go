//Program to read from /proc/self/status file and print the current process info
package main

import (
	"fmt"
	"os"
)

func check(err error) {
	if(err != nil) {
		panic(err)
	}
}

func main() {

	processInfo, err := os.ReadFile("/proc/self/status")
	check(err)

	fmt.Printf("Process Info:\n %s", string(processInfo))
}