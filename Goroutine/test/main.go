package main

import "fmt"

func goroutine(s []string, ch chan string) {
	defer close(ch)
	sum := ""
	for _, v := range s {
		sum += v
		ch <- sum
	}
}

func main() {
	words := []string{"test1", "test2", "test3", "test4"}
	ch := make(chan string)
	go goroutine(words, ch)
	for w := range ch {
		fmt.Println(w)
	}
}
