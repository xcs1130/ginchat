package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式
	//db.AutoMigrate(&models.UserBasic{})
	//db.AutoMigrate(&models.Message{})
	//db.AutoMigrate(&models.GroupBasic{})
	//db.AutoMigrate(&models.Contact{})
	db.AutoMigrate(&models.Community{})

	// // 创建
	// user := &models.UserBasic{}
	// user.Name = "xcs"
	// db.Create(user)

	// // 读取

	// fmt.Println(db.First(user, 1)) // 查询id为1的product
	// //db.First(user, "code = ?", "L1212") // 查询code为l1212的product

	// // 更新 - 更新product的price为2000
	// db.Model(user).Update("PassWord", "12345")

	// // 删除 - 删除product
	// //db.Delete(&product)
}
