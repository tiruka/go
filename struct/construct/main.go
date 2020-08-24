package main

import "fmt"

type Vertex struct {
	x, y int
}

func New(x, y int) *Vertex {
	return &Vertex{x, y}
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func (v *Vertex) area() int {
	return v.x * v.y
}

func main() {
	v := New(3, 4)
	v.Scale(10)
	fmt.Println(v, v.area())
}
