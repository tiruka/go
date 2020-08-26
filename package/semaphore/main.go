package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/semaphore"
)

var s *semaphore.Weighted = semaphore.NewWeighted(1) // 何個同時に走らせるか
var s2 *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	// 順番にlockを取得する。lockの数はsで宣言
	if err := s.Acquire(ctx, 1); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func longProcess2(ctx context.Context) {
	// 1個しか実行されない。lockを取得できないと、終了する
	isAcquired := s2.TryAcquire(1)
	if !isAcquired {
		fmt.Println("Coud not get lock")
		return
	}
	defer s2.Release(1)
	fmt.Println("Wait...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

func main() {
	ctx := context.TODO()
	// go longProcess(ctx)
	// go longProcess(ctx)
	// go longProcess(ctx)
	go longProcess2(ctx)
	go longProcess2(ctx)
	go longProcess2(ctx)
	time.Sleep(5 * time.Second)
}
