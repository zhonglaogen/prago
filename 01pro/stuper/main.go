package main

import (
	"fmt"
	"os"
)

//学员信息管理系统

//1添加学员信息
// 2编辑学员信息
// 3展示所有学员信息

func showMenu() {
	fmt.Println("welcome stu manganer system!")
	fmt.Println("1添加学员")
	fmt.Println("2编辑学员")
	fmt.Println("3展示所有学员信息")
	fmt.Println("4退出")
}

func getInput() *student {

	var (
		id    int
		name  string
		class string
	)

	fmt.Println("输入学员的信息")
	fmt.Scanf("%d %s %s", &id, &name, &class)

	stu := newStudent(id, name, class)
	return stu
}

func main() {
	sm := newStudentMgr()

	for {
		//1.打印系统菜单
		showMenu()
		//2.等待用户选择要执行的选项
		var input int
		fmt.Println("输入要操作的序号:")
		fmt.Scanf("%d\n", &input)
		fmt.Println("输入的是", input)
		//3.执行用户选择的动作
		switch input {
		case 1:
			// 添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			// 编辑学员
			stu := getInput()
			sm.modifyStudent(stu)
		case 3:
			// 展示学员
			sm.showStudent()
		case 4:
			// 退出
			os.Exit(0)
		}
		fmt.Println("\n")
	}

}
