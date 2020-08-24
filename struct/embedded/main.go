package main

import "fmt"

type Vertex struct {
	x, y int
}

func (v *Vertex) Scale(i int) {
	v.x = v.x * i
	v.y = v.y * i
}

func (v *Vertex) area() int {
	return v.x * v.y
}

type Vertex3D struct {
	Vertex
	z int
}

func (v *Vertex3D) Scale3D(i int) {
	v.x = v.x * i
	v.y = v.y * i
	v.z = v.z * i
}

func (v *Vertex3D) area3D() int {
	return v.x * v.y * v.z
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

func main() {
	v := New(3, 4, 5)
	v.Scale(10)
	fmt.Println(v, v.area(), v.area3D())
}
