package main

import "fmt"

func main() {
	a := func() {
		fmt.Print("hello\n")
	}
	a()
	testClosure()
}

//closure
func testClosure() {
	for i := 0; i < 3; i++ {
		defer fmt.Println("a:", i)
		defer func() {
			fmt.Println(i)
		}()
	}
}
