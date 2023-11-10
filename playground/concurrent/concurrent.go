package main

import (
	"fmt"
)

func testRecursion(i int, c chan int) {
	if i == 10 {
		close(c)
		return
	}

	c <- i
	testRecursion(i+1, c)
}

func main() {
	c := make(chan int)
	go testRecursion(0, c)

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Hello, world!")
}
