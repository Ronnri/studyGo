package main

import "fmt"

type student struct {
	id    int //id是唯一的
	name  string
	class string
}

//student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}

}

// 学员管理的类型
type studentMgr struct {
	allStudents []*student
}

func newStudentMgr() *studentMgr {
	return &studentMgr{allStudents: make([]*student, 0, 100)}
}

//1.添加学生
func (s *studentMgr) addStudent(newStudent *student) {
	s.allStudents = append(s.allStudents, newStudent)

}

//2.编辑学生
func (s *studentMgr) modifyStudent(newStudent *student) {
	for k, v := range s.allStudents {
		if newStudent.id == v.id { // 学号相同，找到了对应的学生
			s.allStudents[k] = newStudent // 根据切片，把新学生赋值进来
			return
		}
	}
	//输入的学生不存在
	fmt.Println("查无此人！")

}

//3.展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号：%d，姓名：%s，班级：%s\n", v.id, v.name, v.class)
	}

}
