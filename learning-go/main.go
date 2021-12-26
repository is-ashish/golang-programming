package main

import (
	"fmt"
	"sync"
)

func Run(tasks ...func(*sync.WaitGroup, chan<- error)) error {
	n := len(tasks)
	errCh := make(chan error, n)
	w := &sync.WaitGroup{}

	w.Add(n)
	// Sending into channel
	for _, t := range tasks {
		go func(w *sync.WaitGroup, errCh chan<- error) {
			t(w, errCh)
		}(w, errCh)
	}
	// receiving from channel
	for _, err := range tasks {
		go func(w *sync.WaitGroup, ch <-chan error) {
			data, isChannelOpen := <-ch
			if isChannelOpen {
				fmt.Println(data)
			}
			if err != nil {
				return
			}
		}(w, errCh)

	}
	w.Wait()

	return nil
}
func main() {
	Run()
}
