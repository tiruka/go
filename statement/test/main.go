package main

import "fmt"

func test1() {
	l := []int{100, 300, 23, 11, 2, 4, 6, 4}
	MIN := 10000
	for i := 0; i < len(l); i++ {
		if l[i] < MIN {
			MIN = l[i]
		}
	}
	fmt.Println(MIN)
}

func test2() {
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orage":  80,
		"papaya": 500,
		"kiwi":   50,
	}

	var sum int = 0
	for _, v := range m {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	test1()
	test2()
}
