package main

import (
	"fmt"
	"net/http"

	"github.com/bayurstarcool/bayurGo/app/controllers"
	"github.com/bayurstarcool/bayurGo/route"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	dbUser := "root"
	dbPass := ""
	dbName := "sig"
	dbHost := "localhost"
	dbPort := "3306"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	print("Database connected √\n\n")
	db.LogMode(true)
	return db
}

func main() {
	db := SetupDB()
	router := route.RouteApp(&controllers.AppContext{DB: db})
	http.ListenAndServe(":8080", router)
}
