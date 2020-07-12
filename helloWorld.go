package main

import "fmt"

func main() {
	fmt.Println("Hello World!!!!");
    fmt.Println("Sou um programa em Go, eu existo!");

	var x float64
	x = 20.0
	fmt.Println(x)
	fmt.Printf("x is of %T\n", x)

	var z float64 = 20.0
	y := 42
	fmt.Println(z)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", z)
	fmt.Printf("y is of type %T\n", y)

	var a,b,c = 3, 5, "foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)



}