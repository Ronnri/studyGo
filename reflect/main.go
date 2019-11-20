package main

import (
	"fmt"
	"reflect"
)

//reflect demo

//猜类型， 1.断言   2.反射

//参数类型不确定，空接口类型
//参数类型数量都不确定，可变参数空接口 x... interface{}
func reflectType(x interface{}) {
	vObj := reflect.TypeOf(x)
	fmt.Println(vObj, vObj.Name(), vObj.Kind())
	fmt.Println(vObj)

}

func reflectVal(x interface{}) {
	xVal := reflect.ValueOf(x)
	fmt.Printf("%v\t%T\n", xVal, xVal)

	k := xVal.Kind() // 拿到值对应的类型种类
	switch k {
	case reflect.Float32:
		//把反射取到的值转换成int32类型的变量
		ret := float32(xVal.Float())
		fmt.Printf("%v,%T\n", ret, ret)
	case reflect.Int32:
		ret := int32(xVal.Int())
		fmt.Printf("%v,%T\n", ret, ret)

	}
}

//如果用值：  找不到地址。。。。 找到的只是副本。
func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	//Elem() 用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(23)
	}

}

func main() {
	var a float32 = 123.123
	reflectType(a)
	var b []int
	reflectType(b)

	var aa int32 = 100
	reflectVal(aa)

	reflectSetValue(&aa)
	fmt.Println(aa)

	//匿名结构体
	ret := struct {
		int
		string
	}{
		32,
		"sasdf",
	}
	fmt.Println(ret)
}
