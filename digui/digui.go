package main

import "fmt"

//关键就是找到退出条件
func f(n int64) int64 {
	if n == 1 {
		return 1
	} else {
		return n * f(n-1)
	}
}

func main() {
	fmt.Println("----递归----")
	fmt.Println(f(7))

}
