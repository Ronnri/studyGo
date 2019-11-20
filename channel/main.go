package main

import "fmt"

func main() {

	var ch1 chan int // channel是引用类型，需要初始化后才能使用

	//无缓冲区的通道,又称为同步通道。   有接收者才能发进去 否则会panic
	ch1 = make(chan int)

	//带缓冲区的通道
	ch1 = make(chan int, 1) // 类型 ， 缓冲区
	ch1 <- 10
	x := <-ch1
	fmt.Println(x)

	fmt.Println(len(ch1), cap(ch1))

	close(ch1)

}
