package main

import (
	"fmt"
	"regexp"
)

func main() {

	//s1 := "Z22WhkC9K$N#wRIbBitpQYUFJqd%d*QVz5Ir7j&CLqz3@FOC2x5CD6ZzvhNdk@5Vv1QU^DHbSc&nBs%4!94ngFH0BJ6GxtXg2$T"
	//s1 := "Z22WhkC9K$N#wRIbBitpQYUFJqd%d*QVz5Ir7j&CLqz3@FOC2x5CD6ZzvhNdk@5Vv1QU^DHbSc&nBs%4!"
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	//fmt.Println(reg.FindAllString(s1, -1))

	//reg = regexp.MustCompile(`[a-zA-Z0-9]`)
	//fmt.Println(reg.FindAllString(s1, -1))
	//
	//reg = regexp.MustCompile(`[a-z]`)
	//fmt.Println(reg.FindAllString(s1, -1))
	//
	//reg = regexp.MustCompile(`[A-Z]`)
	//fmt.Println(reg.FindAllString(s1, -1))
	//
	//reg = regexp.MustCompile(`[0-9]`)
	//fmt.Println(reg.FindAllString(s1, -1))
	//
	//fmt.Println(IsEmail("ronnri331@@qq.com"))
	//
	//
	text := `Hello 世界！123 Go.`
	//
	//	// 查找连续的小写字母
	//	reg = regexp.MustCompile(`[a-z]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["ello" "o"]
	//
	//	// 查找连续的非小写字母
	//	reg = regexp.MustCompile(`[^a-z]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["H" " 世界！123 G" "."]
	//
	//	// 查找连续的单词字母
	//	reg = regexp.MustCompile(`[\w]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello" "123" "Go"]
	//
	//	// 查找连续的非单词字母、非空白字符
	//	reg = regexp.MustCompile(`[^\w\s]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["世界！" "."]
	//
	//	// 查找连续的大写字母
	//	reg = regexp.MustCompile(`[[:upper:]]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["H" "G"]
	//
	//	// 查找连续的非 ASCII 字符
	//	reg = regexp.MustCompile(`[[:^ascii:]]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["世界！"]
	//
	//	// 查找连续的标点符号
	//	reg = regexp.MustCompile(`[\pP]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["！" "."]
	//
	//	// 查找连续的非标点符号字符
	//	reg = regexp.MustCompile(`[\PP]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello 世界" "123 Go"]
	//
	//	// 查找连续的汉字
	//	reg = regexp.MustCompile(`[\p{Han}]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["世界"]
	//
	//	// 查找连续的非汉字字符
	//	reg = regexp.MustCompile(`[\P{Han}]+`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello " "！123 Go."]
	//
	//	// 查找 Hello 或 Go
	//	reg = regexp.MustCompile(`Hello|Go`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello" "Go"]
	//
	//	// 查找行首以 H 开头，以空格结尾的字符串
	//	reg = regexp.MustCompile(`^H.*\s`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello 世界！123 "]
	//
	//	// 查找行首以 H 开头，以空白结尾的字符串（非贪婪模式）
	//	reg = regexp.MustCompile(`(?U)^H.*\s`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello "]
	//
	//	// 查找以 hello 开头（忽略大小写），以 Go 结尾的字符串
	//	reg = regexp.MustCompile(`(?i:^hello).*Go`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello 世界！123 Go"]
	//
	//	// 查找 Go.
	//	reg = regexp.MustCompile(`\QGo.\E`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Go."]
	//
	//	// 查找从行首开始，以空格结尾的字符串（非贪婪模式）
	//	reg = regexp.MustCompile(`(?U)^.* `)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello "]
	//
	//	// 查找以空格开头，到行尾结束，中间不包含空格字符串
	//	reg = regexp.MustCompile(` [^ ]*$`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// [" Go."]
	//
	//	// 查找“单词边界”之间的字符串
	//	reg = regexp.MustCompile(`(?U)\b.+\b`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello" " 世界！" "123" " " "Go"]
	//
	//	// 查找连续 1 次到 4 次的非空格字符，并以 o 结尾的字符串
	//	reg = regexp.MustCompile(`[^ ]{1,4}o`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello" "Go"]
	//
	//	// 查找 Hello 或 Go
	//	reg = regexp.MustCompile(`(?:Hell|G)o`)
	//	fmt.Printf("%q\n", reg.FindAllString(text, -1))
	//	// ["Hello" "Go"]
	//
	// 查找 Hello 或 Go，替换为 Hellooo、Gooo
	reg = regexp.MustCompile(`(P:Hell|G)o`)
	fmt.Printf("%q\n", reg.ReplaceAllString(text, "${n}ooo"))
	// "Hellooo 世界！123 Gooo."
	//
	//	// 交换 Hello 和 Go
	//	reg = regexp.MustCompile(`(Hello)(.*)(Go)`)
	//	fmt.Printf("%q\n", reg.ReplaceAllString(text, "$3$2$1"))
	//	// "Go 世界！123 Hello."
	//
	//	// 特殊字符的查找
	//	reg = regexp.MustCompile(`[\f\t\n\r\v\123\x7F\x{10FFFF}\\\^\$\.\*\+\?\{\}\(\)\[\]\|]`)
	//	fmt.Printf("%q\n", reg.ReplaceAllString("\f\t\n\r\v\123\x7F\U0010FFFF\\^$.*+?{}()[]|", "-"))
	//	// "----------------------"

}

func IsEmail(email string) bool {
	if email != "" {
		if isOk, _ := regexp.MatchString("^[_a-z0-9-]+(\\.[_a-z0-9-]+)*@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$", email); isOk {
			return true
		}
	}

	return false
}

func IsPhone(phoneStr string) bool {
	if phoneStr != "" {
		if isOk, _ := regexp.MatchString(`^\([\d]{3}\) [\d]{3}-[\d]{4}$`, phoneStr); isOk {
			return isOk
		}
	}

	return false
}

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
