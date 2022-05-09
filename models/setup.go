package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "admin"
	DB_NAME     = "mydatabase"
	DB_HOST     = "localhost"
	DB_PORT     = 5432
)

// DB set up
func SetUpDB() *gorm.DB {
	// URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// db, err := gorm.Open("postgres", URL)
	dsn := "host=localhost user=postgres password=admin dbname=mydatabase port=5432 sslmode=disable"
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err.Error())
	}
	// defer db.Close() // close the connection when the function returns
	return db
}
