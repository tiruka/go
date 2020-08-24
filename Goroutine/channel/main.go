package main

import "fmt"

func goroutine1(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func goroutine2(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int) // queueのような役割
	go goroutine1(s, c)
	go goroutine2(s, c)
	x := <-c // blockingする。sync.Waitしなくても勝手に待つ。
	fmt.Println(x)
	y := <-c
	fmt.Println(y)
}
