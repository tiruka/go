package main

import "fmt"

func main() {
	var p *int = new(int) // ポインタ型で、メモリ領域だけ確保
	fmt.Println(p, *p)
	*p++
	fmt.Println(*p)

	var p2 *int // ポインタ型は宣言したが、メモリは確保していないのでnil
	fmt.Println(p2)
	// *p++ error

	// pointer型のものはnew, そのほかのものはmakeを使う
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var p3 *int = new(int)
	fmt.Printf("%T\n", p3)

	st := new(struct{})
	fmt.Printf("%T\n", st)
}
