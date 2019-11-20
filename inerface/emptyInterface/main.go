package main

import "fmt"

//空接口
//接口中没有定义任何方法=====>空接口
//任意类型都实现了空接口 ===> 空接口变量可以存储任意值
type xxx interface {
}

func main() {
	var x interface{}
	x = 100
	x = "123"
	x = false
	fmt.Println(x)

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "沙河娜扎"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)

	ret, ok := x.(string)
	if !ok {
		fmt.Println("bushi")
	} else {
		fmt.Println(ret)
	}

	//使用switch进行断言
	switch t := x.(type) {
	case string:
		fmt.Println("s", t)
	case bool:
		fmt.Println("b", t)
	case int:
		fmt.Println("i", t)
	case []string:
		fmt.Println("[]", t)
	default:
		fmt.Println("猜不到了。。。")

	}

}
