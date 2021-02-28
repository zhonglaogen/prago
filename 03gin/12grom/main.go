package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UserInfo struct {
	ID   int
	Name string
	//必须大写,不然没法识别
	Hobby  string
	Gender string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)"+
		"/zlx?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//u1 := UserInfo{1,"zlx","哈哈","男"}
	//db.Create(u1)

	//查询
	var user UserInfo
	db.First(&user)
	fmt.Printf("%#v\n", user)

}
