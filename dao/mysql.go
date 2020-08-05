package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func InitDB() (err error) {
	Db, err = gorm.Open("mysql", "root:1234@/db2?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return
	}
	return Db.DB().Ping()
}

func CloseDB() {
	Db.Close()
}
