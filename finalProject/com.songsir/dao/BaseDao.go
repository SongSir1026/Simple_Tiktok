package dao

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getConnect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3306)/test?parseTime=true"))
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db

}
