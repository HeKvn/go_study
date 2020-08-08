package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

//互斥锁(读写锁，RWMUtex，可并发多读，但只能单写，写的时候不可读)
var mutex sync.Mutex

func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is:", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}
