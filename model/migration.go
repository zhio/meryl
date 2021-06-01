package model

import "fmt"

//执行数据迁移

func migration() {
	//自动迁移模式
	fmt.Print("初始化成功ok")
	_ = DB.AutoMigrate(&User{})
}
