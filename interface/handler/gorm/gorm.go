package gorm

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/tozastation/clasick-core/infrastructure/persistence/model/db"
	"os"
)

type gormHandler struct {}

type IGormHandler interface {
	CreateConnection() (*gorm.DB, error)
}

func NewGormHandler() IGormHandler {
	return &gormHandler{}
}

func (orm *gormHandler) CreateConnection() (*gorm.DB, error) {
	VENDOR := os.Getenv("DB_VENDOR")
	USER     := os.Getenv("USER")
	PASS     := os.Getenv("PASS")
	PROTOCOL := os.Getenv("PROTOCOL")
	DBNAME   := os.Getenv("DBNAME")
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME+ "?parseTime=true"
	DB, err := gorm.Open(VENDOR, CONNECT)
	if err != nil {
		return nil, err
	}
	fmt.Println("Create User Table")
	DB.AutoMigrate(&db.User{})
	return DB, nil
}