//Producer-Consumer Code in Go
package main

import (
	"fmt"
)

func Producer(ch chan int, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("Produced Item %v\n",i)
		ch <- i
	}

	close(ch)
}

func Consumer(ch chan int, done chan bool) {
	for {
		item , more := <-ch
		if more {
		    fmt.Printf("Consumed Item %v\n",item)
		} else {
			done <- true
			return
		}
	}
}

func main() {
	var n int
	fmt.Printf("Enter Number of Items to be produced : ")
	fmt.Scanf("%d\n",&n)
    
	ch := make(chan int, n)
    done := make(chan bool)
	go Producer(ch, n)
	go Consumer(ch, done)

	<-done
}