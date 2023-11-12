package main

import (
    "fmt"
)

func testRecursion(i int, c chan int) {
    defer close(c)
    var testFunc func (int, chan int)
    testFunc = func (i int, c chan int) {
        if i > 10 {
            return
        }
        c <- i
        testFunc(i+1, c)
    }

    testFunc(i, c)
}

func main() {
    c := make(chan int)
    go testRecursion(0, c)

    for i := range c {
        fmt.Println(i)
    }

    fmt.Println("Hello, world!")
}
