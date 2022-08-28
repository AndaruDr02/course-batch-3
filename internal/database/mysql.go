package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDabataseConn() *gorm.DB {
	dsn := "root:P@ssw0rd123!@tcp(localhost:3306)/training?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
