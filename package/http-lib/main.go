package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func req_http() {
	resp, _ := http.Get("http://example.com")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func applied_http() {
	base, _ := url.Parse("http://example.com")
	reference, _ := url.Parse("/test?a=1")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println(endpoint)

	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `wweehh`)
	q := req.URL.Query()
	q.Add("c", "3&da")
	fmt.Println(q)
	fmt.Println(q.Encode())
	req.URL.RawQuery = q.Encode()
	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	// req_http()
	applied_http()
}
