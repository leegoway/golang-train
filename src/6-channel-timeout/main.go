/*
 * 带超时时间的worker，在超时时间内，worker可以工作，处理业务。到达时间后，由timeout条件处理善后工作。
 * 场景： 30秒超时判断任务是否完成。30秒任务超时逻辑，30秒内也可以完成任务。
 */
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(ch chan string) {
	timeout := time.After(2000 * time.Millisecond)
	for {
		select {
		case <-ch:
			fmt.Println("worker done")
		case <-timeout:
			fmt.Println("worker timeout")
			return
		}
	}
}

func pinger(ch chan string) {
	for {
		interval := rand.Intn(1000)
		fmt.Println(interval)
		time.Sleep(time.Duration(interval) * time.Millisecond)
		ch <- "ping"
	}

}

func main() {
	data_cn := make(chan string)
	go pinger(data_cn)
	go worker(data_cn)

	var input string
	fmt.Scan(&input)
}
