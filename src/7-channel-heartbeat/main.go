/*
 * 带hearbeat功能的程序，到达一定时间会处理heartbeat的逻辑，比如发起一个http请求等，哈哈
 */
package main

import (
	"fmt"
	"time"
)

func worker() {
	heartbeat := time.Tick(time.Duration(500) * time.Millisecond)
	for {
		select {
		case <-heartbeat:
			fmt.Println("heart beat")
		}
	}
}

func main() {
	go worker()

	var input string
	fmt.Scan(&input)
}
