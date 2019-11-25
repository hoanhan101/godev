package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool, 1)
	go func() {
		excuteTask(c)
	}()

	v := <-c
	fmt.Println("complete?", v)
}

func excuteTask(complete chan bool) {
	for i := 0; i < 10; i++ {
		fmt.Println("sleeping for a sec", i)
		time.Sleep(1 * time.Second)

		if i == 6 {
			fmt.Println("sending true over")
			complete <- true
			close(complete)
			return
		}
	}
}
