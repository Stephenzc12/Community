package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (err error) {
	host := "localhost"
	port := "3306"
	database := "community"
	username := "root"
	password := "668904"
	charset := "utf8"

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err := CreateTableUser(); err != nil {
		return err
	}
	if err := CreateTableQuestion(); err != nil {
		return err
	}
	if err := CreateTableReply(); err != nil {
		return err
	}

	return
}
