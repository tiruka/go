package main

import "fmt"

func array() {
	var a [2]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	var c []int = []int{100, 200}
	d := append(c, 300)
	fmt.Println(c, d)
}

func slice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:2])
	fmt.Println(n[:])

	n[2] = 100
	fmt.Println(n)
	n = append(n, 100, 200, 300)
	fmt.Println(n)
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)
	board2 := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board2)
}

func main() {
	array()
	slice()
}
