package main

import "time"

func main() {
	c := make(chan int, 1)
	go func() {
		time.Sleep(5 * time.Second)
	}()
	c <- 1
	c <- 1
}
