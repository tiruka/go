package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"time"
)

func timeLib() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format(time.RFC3339))
	fmt.Println(t.Year(), t.Month())
}

func regexLib() {
	match, _ := regexp.MatchString("a[a-z]+e", "apple")
	fmt.Println(match)

	r := regexp.MustCompile("a[a-z]+e")
	ms := r.MatchString("apple")
	fmt.Println(ms)

	s := "/view/test"
	r2 := regexp.MustCompile("^/(edit|save|view)/([a-z]+)$")
	fs := r2.FindString(s)
	fmt.Println(fs)
	fss := r2.FindStringSubmatch("/edit/test")
	fmt.Println(fss)
}

func sortLib() {
	i := []int{5, 3, 7, 9, 1}
	s := []string{"b", "a", "d", "y"}
	p := []struct {
		Name string
		Age  int
	}{
		{"Nancy", 20},
		{"Tom", 33},
		{"Gary", 15},
	}
	fmt.Println(i, s, p)
	fmt.Println("after sorted")
	sort.Ints(i)
	sort.Strings(s)
	sort.Slice(p, func(i, j int) bool { return p[i].Name < p[j].Name })
	fmt.Println(i, s, p)
	sort.Slice(p, func(i, j int) bool { return p[i].Age > p[j].Age }) // Ageでreverse
	fmt.Println(p)
}

func iotaLib() {
	const (
		c1 = iota
		c2
		c3
	)
	fmt.Println(c1, c2, c3)
	const (
		_      = iota
		KB int = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Println(KB, MB, GB, TB)
}

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(3 * time.Second)
	fmt.Println("finish")
	ch <- "result"
}

func contextLib() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)

CTXLOOP:
	for {
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break CTXLOOP
		case <-ch:
			fmt.Println("success")
			break CTXLOOP
		}
	}
	fmt.Println("###")
}

func ioutilLib() {
	content, err := ioutil.ReadFile("main.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	if err := ioutil.WriteFile("ioutil_temp.go", content, 0666); err != nil {
		log.Fatal(err)
	}
}

func main() {
	timeLib()
	regexLib()
	sortLib()
	iotaLib()
	contextLib()
	ioutilLib()
}
