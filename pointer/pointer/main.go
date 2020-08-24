package main

import "fmt"

func one(x int) {
	x = 1
}

func onereal(x *int) {
	// デリファレンス
	*x = 1
}

func arraychange(arr []int) {
	arr = append(arr, 100)
}
func main() {
	var n int = 100
	fmt.Println(n, "n")
	fmt.Println(&n)

	var p *int = &n     // &(アンパサンド)はアドレスをさす。型に*をつかうことで、アドレスを格納させる
	fmt.Println(p, "p") // アドレス
	fmt.Println(*p)     // 実態
	fmt.Println(&p)     // p自体のアドレス

	// ただ値を渡すので、結果は変わらない
	one(n)
	fmt.Println(n, "not change")

	// 実態を渡すので、結果が変わる
	onereal(&n)
	fmt.Println(n, "changed")

	// Pythonでは配列はmutableなので、配列にappendしたら
	// 結果が変わると思ったら、appendがリターンしたものを再代入する都合上、
	// ここでは結果は変わらない。
	// また、配列自体のアドレスはなく、配列の要素ごとにアドレスを取得する
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("before", arr, &arr[0], &arr) // &arr[0]はアドレスが表示されるが, &arrにはでない
	arraychange(arr)
	fmt.Println("after", arr)
}
