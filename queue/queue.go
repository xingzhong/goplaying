package main

import (
	"fmt"
	"math/rand"
	"time"
)

var queue = make(chan int, 5)

func sender() {
	defer close(queue)
	for i := 0; i > -1; i++ {
		queue <- i
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

func worker(idx int) {
	for num := range queue {
		fmt.Printf("[Worker %d] on job %d\n", idx, num)
		time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
	}
}

func stats() {
	for range time.Tick(100 * time.Millisecond) {
		fmt.Printf("queue length = %d\n", len(queue))
	}
}

func main() {
	fmt.Println("use buffered channel as a worker queue")
	for i := 0; i < 3; i++ {
		go worker(i)
	}
	go sender()
	go stats()
	for {
	}
}
