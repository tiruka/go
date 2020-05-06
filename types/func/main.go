package main

import "fmt"

func cal(x int, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y

}

func cal_price(price, item int) (result int) {
	result = price * item
	// return result
	return
}

func main() {
	r1, r2 := cal(10, 20)
	fmt.Println(r1, r2)
	r3 := cal_price(100, 10)
	fmt.Println(r3)
	f := func(x int) {
		fmt.Println("inner function", x)
	}
	f(1)
	func(x int) {
		fmt.Println("inner function", x)
	}(10)
}