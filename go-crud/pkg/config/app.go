package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// setup DB and return it as a connection
var db *gorm.DB

// open the database connection
func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local") // default user of MySQL is root without password

	if err != nil {
		panic(err)
	}
	db = d
}

// return db instance
func GetDB() *gorm.DB {
	return db
}
