package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()
	db.AutoMigrate(&User{})
	//users := []User{
	//	User{Name: "zhangsan", Age: 18, Birthday: time.Now()},
	//	User{
	//		Name:     "libai",
	//		Age:      30,
	//		Birthday: time.Time{},
	//	}}
	//err = db.Create(&users).Error
	//user := User{
	//	Name:     "admin",
	//	Age:      10,
	//	Birthday: time.Time{},
	//}
	//// todo 通过数据的指针来创建
	//err = db.Create(&user).Error // 通过数据的指针来创建
	//if err != nil {
	//	println(err)
	//}
	//db.Select("Name", "Age", "CreatedAt").Create(&user)
	//// omit 忽略
	////db.Omit("Name", "Age", "CreatedAt").Create(&user)
	//
	//first := db.Limit(1).Find(&user)
	//println(first)
	//db.First(&user, 10)
	//
	//println(db.Find(&user).RowsAffected)
	//err = db.Model(&User{}).Where("Age=?", 10).Update("Age=?", 100).Error
	//if err != nil {
	//	println(err)
	//}
	//err = db.Where("id=?", 10).Delete(&User{}).Error
	//if err != nil {
	//	println(err)
	//}

	db.Create(&User{
		Model:    gorm.Model{},
		Name:     "",
		Age:      0,
		Birthday: time.Time{},
	})
	var user User
	db.Where("age", 12).Find(&user)
	db.Model(&User{}).Where("age", 0).Update("age", 13)
	db.Where("id", 0).Delete(&User{})

}
