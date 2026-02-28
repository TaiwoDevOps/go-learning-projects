package main

import "sync"

func main() {
	counter := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	println("Final Counter:", counter)
}
