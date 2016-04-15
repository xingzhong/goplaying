package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("[Example] close a channel will also trigger the select clause")
	fmt.Println("[Example] In fact, close() is equivalent to <- struct{}{}")
	cancel := make(chan struct{})
	timeChannel := time.NewTimer(2 * time.Second)

	go func() {
		select {
		case stuff, ok := <-cancel:
			fmt.Printf("[Running] cancel the process %v %v", stuff, ok)
		case stuff, ok := <-timeChannel.C:
			fmt.Printf("[Running] time is up %v %v", stuff, ok)
		}
	}()

	<-time.After(time.Second)
	close(cancel)
	// cancel <- struct{}{}
	<-time.After(4 * time.Second)
}
