package main

import (
	"log"
	"time"
)

func main() {

	names := []string{"vikash", "kumar"}

	for i, n := range names {
		log.Println(i, n)
		func() {
			log.Println("g", n)
		}()
	}

	time.Sleep(1 * time.Second)

}
