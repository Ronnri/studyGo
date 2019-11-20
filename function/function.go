package main

import (
	"fmt"
	"strings"
)

func d(x ...int) {
	fmt.Println(x)
}

func c() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func c err", err)
		}
	}()
	panic("asdfasdf")
}

//闭包 = 函数 + 外层变量的引用
func b() func() {
	name := "zhangsan"
	return func() {
		fmt.Println("b...hello:", name)
	}
}

func a(name string) func() {
	return func() {
		fmt.Println("a...hello:", name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

//返回类型 两个匿名函数
//每个匿名函数的返回类型为int
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {

	func() {
		fmt.Println("匿名函数")
	}()

	a("ronnri")()
	b()()

	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test")) //test.jpg
	fmt.Println(txtFunc("test")) //test.txt

	fmt.Println(makeSuffixFunc(".png")("123"))

	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
	c()

	fmt.Println("练习题 分金币")
	fmt.Printf("%#v\n", distribution)
	left := dispatchCoin()
	fmt.Println("剩下：", left)

	intSlice := []int{10, 20}

	d(intSlice[0])

}

func dispatchCoin() int {
	for _, name := range users {
		fmt.Printf("%v:\t", name)
		for _, ch := range name {
			fmt.Printf("%c\t", ch)
			switch ch {
			case 'e', 'E':
				{
					coins = coins - 1
					distribution[name] += 1
				}
			case 'i', 'I':
				{
					coins = coins - 2
					distribution[name] += 2
				}
			case 'o', 'O':
				{
					coins = coins - 3
					distribution[name] += 3
				}
			case 'u', 'U':
				{
					coins = coins - 4
					distribution[name] += 4
				}
			}
		}
		fmt.Println("\n")
	}
	return coins
}

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)
