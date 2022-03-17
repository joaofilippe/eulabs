package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() *gorm.DB {
	dbUser := os.Getenv("dbUser")
	dbPass := os.Getenv("dbPass")
	dbPort := os.Getenv("dbPort")
	dbName := os.Getenv("dbName")
	dsn := dbUser + ":" + dbPass + "@tcp(127.0.0.1:" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Algo deu errado na conexão com o banco de dados!")
		fmt.Println(err.Error())
		fmt.Println(dsn)
	}

	return DB
}
