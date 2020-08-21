package main

import "fmt"

func by2num(num int) string {
	if num%2 == 0 {
		return "ok"
	}
	return "no"
}

func main() {
	num := 15
	if num%2 == 0 {
		fmt.Println("even")
	} else if num%5 == 0 {
		fmt.Println("fives")
	} else {
		fmt.Println("odd")
	}
	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("both 10")
	}
	a, b := 11, 10
	if a == 10 || b == 10 {
		fmt.Println("either 10")
	}

	result := by2num(2)
	if result == "ok" {
		fmt.Println("great")
	}

	if result2 := by2num(10); result2 == "ok" {
		fmt.Println("great2")
	}
}
