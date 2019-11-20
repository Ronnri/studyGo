package main

import "fmt"

func modify1(x int) {
	x = 100

}

func modify2(y *int) {
	*y = 100
}
func main() {
	fmt.Println("---指针---")
	a := 10
	b := &a //b表示a的内存地址 ， *b 表示从b的值（也就是a的地址）去读取值，  &b表示查看b的内存地址。
	fmt.Println(a, b, *b, &b)

	//跟C语言一样。
	c := 1
	modify1(c)
	fmt.Println(c)

	modify2(&c)
	fmt.Println(c)

	fmt.Println("---new & make ---")
	var d *int
	d = new(int)
	*d = 10
	fmt.Println(*d)

	var e *int = new(int)
	*e = 12
	fmt.Println(*e)

	var f map[string]int
	f = make(map[string]int, 10)
	f["123"] = 123
	fmt.Println(f)

}
