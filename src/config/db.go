package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	 host := "127.0.0.1"
	 port := "3306"
	 dbname := "go"
	 username := "root"
	 password := ""

	 dsn := username + ":" + password + "@tcp(" + host +":"+port+")/"+dbname+"?charset=utf8&parseTime=true&loc=Local"

	 db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	 if err != nil {
		panic("Tidak dapat terkoneksi dengan database")
	 }

	 return db
}