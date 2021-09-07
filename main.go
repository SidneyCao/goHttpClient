package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	start := time.Now()
	c := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := c.Get("http://107.150.126.50:8080/getTimes?uid=3378735268594028615&sid=6001")
	if err != nil {
		log.Panicf("resp Error %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("io read Error %s", err)
	}
	elapsed := time.Since(start)
	fmt.Printf("response: %s, elapsed: %s\n", string(body), elapsed)

}
