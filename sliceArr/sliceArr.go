package main

import (
	"fmt"
)

func main() {
	a := make([]string, 2, 10)
	fmt.Println(a)
	a[0] = "123"
	fmt.Println(a)
	a = append(a, "nihaoshijie")
	fmt.Println(a)

	s := []int{
		1,
		2,
		4,
		8,
	}

	for k, v := range s {
		fmt.Println(k, v)
	}

	s = append(s[:2], s[3:]...) // delete element
	fmt.Println(s)

	scoreMap := make(map[string]int, 10) //map 是无序的
	scoreMap["a"] = 1
	scoreMap["b"] = 2
	fmt.Println(scoreMap)
	fmt.Printf("%#v \n", scoreMap)

	userInfo := map[string]string{
		"a": "123",
		"b": "456",
	}
	fmt.Println(userInfo)

	v, ok := userInfo["a"]
	if ok {
		println("hhh" + v)
	} else {
		println("www")
	}

	////顺序遍历map
	//scoreMap1 := make(map[string]int,100)
	//
	//for i:= 0; i < 100 ; i ++ {
	//	key := fmt.Sprintf("std%02d",i)
	//	val := rand.Intn(100)
	//	scoreMap1[key] = val
	//}
	////fmt.Println(scoreMap1)//理论上是无序的
	//
	////将key取出放入切片中
	//keys := make([]string,100)
	//for key:= range scoreMap1{
	//	keys = append(keys,key)
	//}
	////fmt.Println(keys)
	//
	//sort.Strings(keys)
	////fmt.Println(keys)
	//
	//for _,key := range keys{
	//	fmt.Println(key,scoreMap1[key])
	//}

	//元素类型为map的切片
	mapSlice := make([]map[string]int, 10) //此时切片初始化了，可以使用，但是内部的map没有初始化，无法使用
	fmt.Println(mapSlice)
	for k, v := range mapSlice {
		fmt.Println(k, v)
	}
	fmt.Println(mapSlice[0])

	mapSlice[0] = make(map[string]int, 5) //对map进行初始化,此时，可以正常使用
	fmt.Printf("%T \t %T\n", mapSlice[0], mapSlice)

	fmt.Println("\n\n\n\n//值为切片的map")
	//值为切片的map
	sliceMap := make(map[string][]string, 10) // map初始化
	//fmt.Println(sliceMap)
	key := "city"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2) // 切片初始化
	}
	value = append(value, "tokyo", "hongkong")
	sliceMap[key] = value
	fmt.Println(sliceMap)

}
