package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello(i int) {
	fmt.Printf("hello!!!This is gorountine[%d]\n", i)

	wg.Done()
}
func main() {
	startTime := time.Now()
	runtime.GOMAXPROCS(16)

	//wg.Add(1) //计数，  线程+1
	for i := 0; i < 1000000; i++ {
		wg.Add(1)   //执行一次+1
		go hello(i) //开启了一个独立的goroutine
	}

	fmt.Println("hello main")
	wg.Wait() //等所有线程结束后才退出。
	endTime := time.Now().Sub(startTime)

	fmt.Println("耗时----->", endTime)
}
