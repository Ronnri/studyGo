package main

import "fmt"

//两个goroutine
// 1. 生成0-100的数字发送到ch1中
// 2. 从ch1中取出数据并计算它的平方，把结果发送到ch2中。

//单向通道
//只写  chan <-    	只能往通道里面发东西
//只读  <-chan		只能从通道里面取东西
func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1, ch2 chan int) {
	//从通道中循环取值  方式1
	for {
		tmp, ok := <-ch1
		if ok {
			ch2 <- tmp * tmp
			//fmt.Printf("get data from ch1, value= %v",tmp)
		} else {
			break
		}
	}
	close(ch2)

}

func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	go f1(ch1)
	go f2(ch1, ch2)

	//从通道中循环取值    方式2
	for ret := range ch2 {
		fmt.Println(ret)
	}

}
