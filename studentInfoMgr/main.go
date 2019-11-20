package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

//需求:
// 1.添加学员信息
// 2.编辑学员信息
// 3.展示所有学员信息

func showMenu() {
	fmt.Println("welcome~")
	fmt.Println("1.添加学员信息")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出系统")

}

//获取用户输入的学员信息
func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("请按要求输入学员信息")
	fmt.Printf("请输入学员的id：")
	fmt.Scanf("%d\n", &id)
	fmt.Printf("请输入学员的name：")
	fmt.Scanf("%s\n", &name)
	fmt.Printf("请输入学员的class：")
	fmt.Scanf("%s\n", &class)

	stu := newStudent(id, name, class) // 构造一个学生
	return stu
}

func main() {
	sm := newStudentMgr()
	for {
		//1. 打印系统菜单
		showMenu()
		//2.等待用户选择
		var ch int
		fmt.Printf("请输入你的选择：")
		fmt.Scanf("%d\n", &ch)
		//3.执行用户的选择
		switch ch {
		case 1:
			//添加学员
			std := getInput()
			sm.addStudent(std)
		case 2:
			//编辑学员
			std := getInput()
			sm.modifyStudent(std)
		case 3:
			//展示所有学员
			sm.showStudent()
		case 4:
			//退出
			os.Exit(0)

		}
	}
}
