package main

import "fmt"

//为什么需要接口
type dog struct {
	name string
}

type cat struct {
	name string
}

type peroson struct {
	name string
}

func (d dog) say() {
	fmt.Println("wangwangwang~")

}

func (c cat) say() {
	fmt.Println("miaomiaomiao~")

}

func (p peroson) say() {
	fmt.Println("wuwuwuwu")
}

//打
//参数的类型？？？

func da(arg sayer) {
	arg.say() // 执行say方法， 不管对象是什么

}

//接口不管是什么类型，只管要实现的方法
//定义一个类型，一个抽象的类型，只要实现了say()这个方法的类型都可以成为sayer类型。
type sayer interface {
	say()
}

func main() {
	c1 := cat{}
	da(c1)

	p1 := peroson{name: "zhangsan"}

	da(p1)

	var s sayer
	c2 := cat{}
	s = c2
	p2 := peroson{name: "lisi"}
	s = p2
	fmt.Println(s)

}
