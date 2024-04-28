package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	// ReturnReceiveOnlyChannelsAsResults()
	// PassSendOnlyChannelsAsArguments()
	// TheFirstResponseWins()
	OneToOneNotificationBySendingAValueToAChannel()

	elapsedTime := time.Since(startTime)
	fmt.Println("Execution time:", elapsedTime)
}
