/*
 * 从[0,10]中生成随机数字，但数字为5时就不发送ping了。然后select就获取不到chan值，于是就走timeout了。
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pinger(c chan string) {
	timeout := 5
	alive := true
	for alive {
		timeout = rand.Intn(10)
		if timeout == 5 {
			alive = false
		}
		fmt.Print(timeout)
		c <- "ping"
		time.Sleep(time.Duration(900) * time.Millisecond)
	}
}

func main() {
	var ch chan string = make(chan string)
	timeout := false

	go pinger(ch)

	for timeout != true {
		time.Sleep(time.Duration(1000) * time.Millisecond)
		select {
		case <-ch:
			fmt.Println("pong")
		default:
			fmt.Println("timeout")
			timeout = true
		}
	}
}
