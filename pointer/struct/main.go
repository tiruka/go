package main

import "fmt"

// Vertex :Capital caseは外部からアクセス可能、小文字はprivateになる
type Vertex struct {
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	v.X = 1000
	// こちらと同じだが、書かない(*v).X = 1000
}

func main() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)

	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3)

	// v4, v5は同じ意味
	v4 := Vertex{}
	fmt.Println(v4)

	var v5 Vertex
	fmt.Println(v5)

	// v6, v7は同じ意味でpointerを返す
	v6 := new(Vertex)
	fmt.Println(v6)
	fmt.Printf("%T %v\n", v6, v6)

	v7 := &Vertex{}
	fmt.Println(v7)
	fmt.Printf("%T %v\n", v7, v7)

	a := Vertex{1, 2, "test"}
	changeVertex(a)
	fmt.Println(a, "func")
	b := &Vertex{1, 2, "test"}
	changeVertex2(b)
	fmt.Println(b, *b)
}
