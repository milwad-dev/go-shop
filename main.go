package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func main() {
	_, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go-shop?charset=utf8"))
	if err != nil {
		panic("failed to connect database")
	}

	// Routes
	r := mux.NewRouter()

	fmt.Println("Application run on port :8080")
	if err = http.ListenAndServe(":8080", r); err != nil {
		panic("Error when server application: " + err.Error())
	}
}
