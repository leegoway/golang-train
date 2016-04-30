package main

import (
	"fmt"
	"time"
)

func pinger(tickerNumber int, c chan string) {
	for i := 0; ; i++ {
		fmt.Println("before ping", i)
		c <- fmt.Sprint(tickerNumber, "ping")
		fmt.Println("after ping", i)
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println("handle", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string, 2)

	go pinger(1, c)
    go printer(c)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Got message:", input)
}