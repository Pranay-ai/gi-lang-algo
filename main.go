package main

import (
	"sync"
	"time"

	"algo.com/v2/tokenBucket"
)

func main() {

	tokenBucket := tokenBucket.NewTokenBucket(10, 0.5)

	var wg sync.WaitGroup

	for i := 0; i <= 40; i = i + 2 {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if tokenBucket.AllowRequest() {
				println("Request allowed-", id)
			} else {
				println("Request denied-", id)
			}
		}(i)

		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if tokenBucket.AllowRequest() {
				println("Request allowed-", id)
			} else {
				println("Request denied-", id)
			}
		}(i + 1)

		time.Sleep(300 * time.Millisecond)
	}

	wg.Wait()

}
