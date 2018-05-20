package main

import "fmt"

func main() {
	v := 42
	fmt.Printf("v is of type %T\n", v)

	i := 3.142
	fmt.Printf("i is of type %T\n", i)

	g := 0.867 + 0.5i
	fmt.Printf("g is of type %T\n", g)
}