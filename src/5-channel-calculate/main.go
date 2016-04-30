/*
 * channel 为函数返回值，通过返回只读的channel让其他进程获取消息
 */
package main

import (
	"fmt"
)

func calculate() <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", "calculate", i)
		}
	}()
	return c
}

func main() {

	c := calculate()

	for i := 0; i < 5; i++ {
		fmt.Println("You say:", <-c)
	}
	fmt.Println("OK! I am leaving")
}
