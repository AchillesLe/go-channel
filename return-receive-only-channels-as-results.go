package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 5)
		r <- rand.Int31n(100)
	}()

	return r
}

func ReturnReceiveOnlyChannelsAsResults() {
	a, b := 3, longTimeRequest()
	total := sumSquares(int32(a), <-b)
	fmt.Println("total = ", total)
}
