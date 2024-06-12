package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		ResponseHeaderTimeout: time.Hour,
		MaxConnsPerHost:       99999,
		DisableKeepAlives:     true,
	}

	myClient := &http.Client{Transport: tr}
	for i := 0; i < 60000; i++ {
		go func(n int) {
			_, err := myClient.Get("http://localhost:8080")
			if err != nil {
				fmt.Printf("%d: %s\n", n, err.Error())
			}
		}(i)
		time.Sleep(1 * time.Millisecond)
		if i%5000 == 0 || i == 59998 {
			fmt.Println("Sleeping for 1 sec")
			time.Sleep(1 * time.Second)
		}
	}
	time.Sleep(time.Hour)
}
