package main

import "fmt"

// var is valid without function
var i int = 10
var f64 float64 = 1.2
var s string = "test"
var t, f bool = true, false
var (
	a int = 1
	b int = 2
	c int
)

func foo() {
	// short variable declaration
	xi := 1
	xi = 2
	xf64 := 1.2
	xs := "short hand"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)
	// check type
	fmt.Printf("%T\n", xf64)
}
func main() {
	foo()
	fmt.Println(i, f64, s, t, f)
	fmt.Println(a, b, c)
}
