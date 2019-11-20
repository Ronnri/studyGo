package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type student struct {
	name string `json:"name" ini:"STU_NAME"`
	age  int    `json:"age" ini:"STU_AGE"`
}

func main() {

	stu := student{
		name: "ZHANGSAN",
		age:  12,
	}

	//通过反射去获取结构体中所有的字段信息
	t := reflect.TypeOf(stu)
	fmt.Printf("name:%v\tkind:%v\n", t.Name(), t.Kind())

	//遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		//i就是结构体字段的索引
		//根据i 去取字段
		filedObj := t.Field(i)
		fmt.Printf("name:%v,type:%v,tag:%v\n", filedObj.Name, filedObj.Type, filedObj.Tag)
		fmt.Println(filedObj.Tag.Get("json"), filedObj.Tag.Get("ini"))

	}
	//根据名字取结构体字段
	fmt.Println("根据名字取结构体字段")
	filedObj2, ok := t.FieldByName("age")
	if ok {
		fmt.Printf("name:%v,type:%v,tag:%v", filedObj2.Name, filedObj2.Type, filedObj2.Tag)

	} else {

	}

}
