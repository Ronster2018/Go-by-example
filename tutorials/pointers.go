package main

import "fmt"

func main() {
	x := 15
	a := &x // the memory address
	fmt.Println(a)
	fmt.Println(*a) // Reads the content in the memory address

	*a = 5 // Changing the value of the above memory address
	fmt.Println(x)

}
