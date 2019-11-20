package main

import (
	"fmt"
	"math"
)

var (
	a int     = 1
	b float32 = 2.2
	c bool    = false
	d string  = "helloworld你好世界"
	e         = [...]int{1, 2, 3}
)

func foo() (int, string) {
	return 10, "hello"
}

func printType() {
	fmt.Printf("%T\t(%v)\n", a, a)
	fmt.Printf("%T\t(%v)\n", b, b)
	fmt.Printf("%T\t(%v)\n", c, c)
	fmt.Printf("%T\t(%v)\n", d, d)
	fmt.Printf("%T\t(%v)\n", e, e)
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		// forloop2:
		for j := 0; j < 5; j++ {
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}

func cfb() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("(%v)*(%v) = %v\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

func bianli() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	// fmt.Println(a)
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

func qiepian() {
	//切片再切片
	a := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("a:%v type:%T len:%d  cap:%d\n", a, a, len(a), cap(a))
	b := a[1:3]
	fmt.Printf("b:%v type:%T len:%d  cap:%d\n", b, b, len(b), cap(b))
	c := b[1:5]
	fmt.Printf("c:%v type:%T len:%d  cap:%d\n", c, c, len(c), cap(c))
}

func qiepiantest() {
	var a = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}

func main() {
	fmt.Println("hello world")

	x, _ := foo() //匿名变量
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	fmt.Printf("%f", math.Pi)
	fmt.Printf("%.2f\n\n\n", math.Pi)

	printType()

	counter := 0
	for _, r := range d {
		if r > 255 {
			counter++
			fmt.Printf("%c", r)
		}
	}
	fmt.Println(counter)

	// continueDemo()
	cfb()

	for index, value := range e {
		fmt.Println(index, value)
	}

	bianli()
	qiepian()
	qiepiantest()

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice = append(newSlice, 60)
	fmt.Println(slice, newSlice)
	slice1 := []int{10, 20, 30, 40}
	newSlice1 := append(slice1, 50)
	fmt.Println(slice1, newSlice1)

	var arr = [...]int{1, 2, 3, 4, 5}
	arrSlice := arr[1:4]
	arrSlice = append(arrSlice, 888, 999)
	fmt.Println(arrSlice, arr)

	mapSlice := make([]map[string]int, 10) //元素类型为map的切片
	fmt.Println(mapSlice)
	for k, v := range mapSlice {
		fmt.Println(k, v)
	}
	fmt.Println(mapSlice[0])

}
