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
type AdminUsers struct {
	Admin bool
}

// 表名映射 但是不是动态的
func (u User) TableName() string {
	return "sss"
}

func main() {
	dsn := "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	var db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db = db.Debug()
	db.AutoMigrate(&Good{})
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

	//db.Create(&User{
	//	Model:    gorm.Model{},
	//	Name:     "",
	//	Age:      0,
	//	Birthday: time.Time{},
	//})
	//var user User
	//db.Where("age", 12).Find(&user)
	//db.Model(&User{}).Where("age", 0).Update("age", 13)
	//db.Where("id", 0).Delete(&User{})
	////指定数据库表名
	//err = db.AutoMigrate(&AdminUsers{})
	//if err != nil {
	//	return
	//}
	//db.Table("admin_users").Create(&AdminUsers{Admin: true})

	//Save保存所有字段，用于单个记录的全字段更新
	db.Order("id desc").Limit(2).Offset(3).Find(&User{})
	type Result struct {
		Name string
		Age  int
	}
	var result Result
	db.Table("users").Select("name", "age").Where("name = ?", "Antonio").Scan(&result)

	//会话 事务
	err = db.Session(&gorm.Session{}).Transaction(func(*gorm.DB) error {
		return nil
	})
	if err != nil {
		return
	}
	//钩子
	//// 开始事务
	//BeforeSave
	//BeforeCreate
	//// 关联前的 save
	//// 插入记录至 db
	//// 关联后的 save
	//AfterCreate
	//AfterSave
	//// 提交或回滚事务
	good := Good{
		Id:   2,
		Name: "zhangsan",
	}
	err = db.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(&good).Error
		return nil
	})
	if err != nil {
		println(err)
	}
	//原生sql语句不执行钩子函数
	sql := "insert into......."
	db.Raw(sql, "")
}

// 数据模型
type Good struct {
	Id   int
	Name string
}

func (*Good) AfterCreate(tx *gorm.DB) error {
	print("-----after create ......")
	return nil
}
func (*Good) BeforeCreate(*gorm.DB) error {
	print("-----before create ......")
	return nil
}
