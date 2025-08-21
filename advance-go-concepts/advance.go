package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ------------------ Interfaces ------------------
type Reader interface {
	Read() string
}

type File struct{}

func (f File) Read() string { return "file content" }

// ------------------ Pointers ------------------
func increment(x *int) {
	*x++
}

// ------------------ Goroutines ------------------
func goroutineDemo() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock() // prevent race
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println("Final counter:", counter)

	// Channels with close and select
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println("channel value:", v)
	}

	ch1 := make(chan string, 1)
	go func() { ch1 <- "message" }()

	select {
	case msg := <-ch1:
		fmt.Println("select got:", msg)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}

// ------------------ Context ------------------
func contextDemo() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("finished work")
	case <-ctx.Done():
		fmt.Println("context cancelled:", ctx.Err())
	}
}

// ------------------ Zero Values ------------------
func zeroValuesDemo() {
	var s string
	var n int
	var p *int
	fmt.Printf("Zero values -> string: '%s', int: %d, pointer: %v\n", s, n, p)
}

// ------------------ Patterns ------------------
// Worker pool pattern
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		results <- j * 2
		fmt.Printf("worker %d processed job %d\n", id, j)
	}
}

func workerPoolDemo() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println("result:", r)
	}
}

// Repository pattern

type User struct {
	ID   int
	Name string
}

type UserRepo interface{ Find(id int) User }

type SQLUserRepo struct{}

func (r SQLUserRepo) Find(id int) User { return User{ID: id, Name: "User"} }

// ------------------ Primitive & Composite ------------------
type Person struct {
	Name string
	Age  int
}

// ------------------ Architecture (Layered Example) ------------------
// Handler -> Service -> Repository

// Service

type UserService struct{ repo UserRepo }

func (s UserService) GetUser(id int) User {
	return s.repo.Find(id)
}

// Handler

func handler() {
	service := UserService{repo: SQLUserRepo{}}
	user := service.GetUser(42)
	fmt.Println("Handler got user:", user)
}

// ------------------ Main ------------------
func main() {
	// Interface demo
	var r Reader = File{}
	fmt.Println("Interface Read:", r.Read())

	// Pointer demo
	a := 5
	increment(&a)
	fmt.Println("Pointer increment:", a)

	// Goroutines demo
	goroutineDemo()

	// Context demo
	contextDemo()

	// Zero values demo
	zeroValuesDemo()

	// Worker pool demo
	workerPoolDemo()

	// Repository & layered architecture demo
	handler()
}
