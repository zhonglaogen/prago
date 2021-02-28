package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string `gorm:"default:'哈喽'"`
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
//唯一表名
func (User) TableName() string {
	return "profiles"
}

type UserInfo struct {
	ID   int
	Name string
	//必须大写,不然没法识别
	Hobby  string
	Gender string
}

func main() {
	//默认表明的前缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "prefix_" + defaultTableName;
	}

	//desc 表名  查看表结构
	db, err := gorm.Open("mysql", "root:root@(localhost:3306)"+
		"/zlx?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//禁用复数
	//db.SingularTable(true)

	//创建表, 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{})

	//使用user结构体创建haha的表
	//db.Table("haha").CreateTable(&UserInfo{})

	//创建数据行

	//u1 := UserInfo{1,"zlx","吃","男"}
	//db.Create(&u1)
	//检查是否存在
	//db.NewRecord(u1);

	//查询
	var u1 UserInfo
	db.First(&u1)
	fmt.Printf("u:%#v\n", u1)

	//	更新
	db.Model(&u1).Update("hobby", "池")
	fmt.Printf("u:%#v\n", u1)
	//	删除
	db.Delete(&u1)
	fmt.Printf("u:%#v\n", u1)
}
