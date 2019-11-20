package main

import "fmt"

type myInt int

func main() {
	var n myInt
	n = 100
	fmt.Printf("%T,%v", n, n)

}
