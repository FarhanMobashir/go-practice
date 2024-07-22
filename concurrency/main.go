package main

import (
	"fmt"
	"sync"
	"time"
)

// question - Write a program that runs two goroutines and prints their outputs.
func goroutinePrint() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 1")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Goroutine 2")
	}()

	wg.Wait()
}

// question 2 - Create a producer-consumer problem using channels.
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Received:", num)
	}
}

// question 3 - Implement a program that uses the select statement to handle multiple channels.
func selectChannelDemo() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)

		}
	}
}

// question 4 - Write a program that uses a mutex to synchronize access to shared data.
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value

}

// question 5 - Implement a worker pool to process tasks concurrently.
type Task func()
type WorkerPool struct {
	tasks chan Task
	wg    sync.WaitGroup
}

func NewWorkerPool(numWorker int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan Task),
	}

	for i := 0; i < numWorker; i++ {
		go pool.worker()
	}
	return pool
}

func (p *WorkerPool) worker() {
	for task := range p.tasks {
		task()
		p.wg.Done()
	}
}

func (p *WorkerPool) Submit(task Task) {
	p.wg.Add(1)
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func (p *WorkerPool) Close() {
	close(p.tasks)
	p.Wait()
}

func main() {
	fmt.Println("Hello from concurrency")
	//goroutinePrint()
	// ch := make(chan int)
	// go producer(ch)
	// consumer(ch)
	//selectChannelDemo()

	//counter := SafeCounter{}

	//var wg sync.WaitGroup
	//numGoroutines := 5

	// wg.Add(numGoroutines)
	// for i := 0; i < numGoroutines; i++ {
	// go func() {
	// defer wg.Done()
	// for j := 0; j < 1000; j++ {
	//	counter.Increment()
	// }
	//	}()
	// }

	//wg.Wait()
	// fmt.Printf("Final counter value: %d\n", counter.Value())

	// Worker pool demo
	pool := NewWorkerPool(3)

	for i := 0; i < 10; i++ {
		taskID := i
		pool.Submit(func() {
			fmt.Printf("Processing task %d\n", taskID)
			time.Sleep(time.Second)
		})
	}
	pool.Close()
}
