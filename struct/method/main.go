package main

import "fmt"

type Vertex struct {
	X, Y int
}

// 引数に対象のstructを指定するのが、pythonでいうselfみたいに
// class内部に持たせることなのかな。全然違った。下の方がその方法
func area(v Vertex) int {
	return v.X * v.Y
}

func (v Vertex) area() int {
	return v.X * v.Y
}

func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(area(v))
	fmt.Println(v.area())
	v.Scale(10)
	fmt.Println(v)
}
