// Go program to illustrate
// the concept of Goroutine
package main

import (
	"fmt"
	"time"
)

func displayRoutines(str1 string) {
	for z := 0; z < 2; z++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str1)
	}
}
func myfun(mychnl chan string) {

	for v := 0; v < 4; v++ {
		mychnl <- "Ashish Jain"
	}
	close(mychnl)
}

func main() {
	// goroutine
	// Calling Goroutine
	// go displayRoutines("Welcome")
	// Calling normal function
	// displayRoutines("GeeksforGeeks")
	// =================================================================
	// Creating a channel
	c := make(chan string)
	// calling Goroutine
	go myfun(c)
	// When the value of ok is
	// set to true means the
	// channel is open and it
	// can send or receive data
	// When the value of ok is set to
	// false means the channel is closed
	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
	mychnl := make(chan string, 4)

	// Anonymous goroutine
	go func() {
		mychnl <- "GFG"
		mychnl <- "gfg"
		mychnl <- "Geeks"
		mychnl <- "GeeksforGeeks"
		fmt.Println("Length of the channel is: ", len(mychnl))
		fmt.Println("Capacity of the channel is: ", cap(mychnl))

		close(mychnl)
		fmt.Println("Length of the channel is: ", len(mychnl))
		fmt.Println("Capacity of the channel is: ", cap(mychnl))

	}()

	// Using for loop
	for res := range mychnl {
		fmt.Println(res)
	}
}
