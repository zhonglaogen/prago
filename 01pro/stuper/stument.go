package main

import "fmt"

type student struct {
	id    int // 学号是唯一的
	name  string
	class string
}

// 学员管理的结构体

func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudent []*student
}

func newStudentMgr() *studentMgr {

	return &studentMgr{
		allStudent: make([]*student, 0, 100),
	}
}

// 1添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudent = append(s.allStudent, newStu)
}

// 2编辑学生

func (s *studentMgr) modifyStudent(newStu *student) {
	for i, v := range s.allStudent {
		if newStu.id == v.id {
			s.allStudent[i] = newStu
			return
		}
	}
	fmt.Println("不存在此学生")
}

// 3展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudent {
		fmt.Printf("学号: %d 姓名: %s 班级: %s\n", v.id, v.name, v.class)
	}
}
