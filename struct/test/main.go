package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Plus() int {
	r := v.X + v.Y
	return r
}

func (v Vertex) String() string {
	return fmt.Sprintf("x is %v! y is %v!", v.X, v.Y)
}

func test1() {
	v := Vertex{3, 4}
	fmt.Println(v.Plus())
}

func test2() {
	v := Vertex{3, 4}
	fmt.Println(v)
}

func main() {
	test1()
	test2()
}
