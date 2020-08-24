package main

// sync.WaitGroupを作って、Waitさせるのは、Pythonでいうところの
// threadを作って、joinさせるのに似ている
// ちなみに、wgを使わないと、goroutine()の実行が終わる前に（スレッドを作っている最中に）、
// main()が終了してしまう。(time.Sleepが書いてあるのは、wgなしで便宜的に表示させるため)

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s, &wg)
	}
}

func normal(s string) {
	for i := 0; i < 5; i++ {
		// time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("world", &wg)
	normal("hello")
	wg.Wait()
}
