package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type test struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:214711@(127.0.0.1:3306)/bloglist?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&test{})
	//u1 := test{ID: 1, Name: "claire", Age: 19}
	//db.Create(&u1)
	var u test
	db.First(&u) //查询表中第一条数据
	fmt.Println("u:%#v\n", u)
	//更新
	db.Model(&u).Update("Age", 20)
	//删除
	db.Delete(&u)
}
