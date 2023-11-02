package main

import (
	"fmt"
	"time"
)

func main() {
	// write a loop that will run 100000 times and print the time
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		fmt.Printf("Index value is %d\n", i)

	}

	endTime := time.Now()
	takenTime := endTime.Sub(startTime)
	fmt.Printf("\nTime taken is %s\n", takenTime)
}

/***
Index value is 99999
Time taken is 258.083083ms
**/
