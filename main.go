package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go-shop?charset=utf8"))
	if err != nil {
		panic("failed to connect database")
	}
}
