package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go incrementor("bar:")
	go incrementor("foo:")
	wg.Wait()
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		mutex.Lock()
		x := counter
		x++
		counter = x
		mutex.Unlock()
		fmt.Println(s, i, "counter:", counter)
	}
	wg.Done()
}
