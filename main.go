package main

import (
	"fmt"
	"github.com/milwad-dev/go-shop/internal/models"
	"github.com/milwad-dev/go-shop/internal/routers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	// Connect DB and migrate
	db := initDB()

	// Routes
	r := routers.InitRouters(db)

	fmt.Println("Application run on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Error when server application: " + err.Error())
	}
}

func initDB() *gorm.DB {
	dialector := mysql.Open("root:@tcp(127.0.0.1:3306)/go-shop?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrations
	db.AutoMigrate(&models.User{})

	return db
}
