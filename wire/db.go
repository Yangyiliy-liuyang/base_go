package wire

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, _ := gorm.Open(mysql.Open("xxx"))
	return db
}
