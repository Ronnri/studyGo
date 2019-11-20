package main

import (
	"fmt"
	"sync"
)

//多个goroutine并发操作全局变量x

var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex //定义一个互斥锁。===> 完全互斥
)

func add() {
	for i := 0; i < 50000; i++ {
		//取值   加1  返回   存在竞争
		//x ++
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁

	}
	wg.Done()

}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
