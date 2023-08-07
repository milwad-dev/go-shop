package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/milwad-dev/go-shop/internal/handlers"
	"github.com/milwad-dev/go-shop/internal/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	// Connect DB and migrate
	db := initDB()

	// Routes
	r := initRouters(db)

	fmt.Println("Application run on port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("Error when server application: " + err.Error())
	}
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go-shop?charset=utf8"))
	if err != nil {
		panic("failed to connect database")
	}

	// Migrations
	db.AutoMigrate(&models.User{})

	return db
}

func initRouters(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()

	// Users routes
	userRepo := handlers.SetRepositories(db)
	r.HandleFunc("/users", userRepo.UserIndex)
	r.HandleFunc("/users/{id}", userRepo.UserShow)

	// Return router instance
	return r
}
