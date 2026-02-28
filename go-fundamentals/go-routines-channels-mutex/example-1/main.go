package main

import (
	"fmt"
	"sync"
	"time"
)

var usersRes = make(chan []string, 10)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println(message)
}

func main() {

	usersConnected := 10
	var wg sync.WaitGroup

	for users := range usersConnected {
		wg.Add(1)
		go sayHello(fmt.Sprintf("User %d", users+1), 2*time.Second, &wg)
	}

	// for res := range usersRes {
	// 	fmt.Println(res)
	// }

	wg.Wait()
}
