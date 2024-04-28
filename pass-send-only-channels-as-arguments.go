package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest1(r chan<- int32) {
	time.Sleep(time.Second * 3)
	r <- rand.Int31n(100)
}

func PassSendOnlyChannelsAsArguments() {
	results := make(chan int32, 2)
	go longTimeRequest1(results)
	go longTimeRequest1(results)

	total := sumSquares(<-results, <-results)
	fmt.Println("total = ", total)
}
