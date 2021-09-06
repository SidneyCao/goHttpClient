package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	for j := 0; j < 1; j++ {
		for i := 0; i < 100; i++ {
			go request()
		}
	}
	time.Sleep(10 * time.Second)
}

func request() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("http://107.150.126.50:8080/getTimes?uid=3378735268594028615&sid=6001")
	if err != nil {
		fmt.Printf("Error %s", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}
