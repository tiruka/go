package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}
	sum := 1
	for sum < 10 {
		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum, "sum")

	arr := []string{"python", "go", "react"}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	for _, v := range arr {
		fmt.Println(v)
	}
	for j, v := range arr {
		fmt.Println(j, v)
	}

	m := map[string]int{"apple": 3, "orange": 5}
	for k, val := range m {
		fmt.Println(k, val)
	}
}
