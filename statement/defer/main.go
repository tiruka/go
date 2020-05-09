package main

import "fmt"

func some() {
	defer fmt.Println("Hello World") // after completed this func, deal this defer
	fmt.Println("Hello Go")
}

func main() {
	// defer some()
	// fmt.Println("Main func")
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3) // stacking process
	fmt.Println("success")
}
