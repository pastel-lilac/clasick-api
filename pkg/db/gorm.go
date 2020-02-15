package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)

var (
	_useDB = os.Getenv("USE_DB")
	_connectionString = os.Getenv("CONNECTION_STRING")
)

func init() {
	log.Printf("info: read database environment variable")
	log.Printf("info: USE_DB -> %s", _useDB)
	log.Printf("info: CONNECTION_STRING -> %s", _connectionString)
}

type (
	IGormDriver interface {
		GetDB() *gorm.DB
		Migrate(args ...interface{}) (err error)
	}
	GormDriver struct {
		DB *gorm.DB
	}
)

func NewGormDriver() IGormDriver {
	driver := &GormDriver{}
	if err := driver.Open(); err != nil {
		log.Fatalf("error: could not open database %v", err)
	}
	return driver
}

func (gd *GormDriver) GetDB() *gorm.DB {
	return gd.DB
}

func (gd *GormDriver) Open() (err error) {
	gd.DB, err = gorm.Open(_useDB, _connectionString)
	if err != nil {
		return err
	}
	return nil
}

func (gd *GormDriver) Migrate(args ...interface{}) (err error) {
	if err := gd.GetDB().AutoMigrate(args).Error; err != nil {
		return err
	}
	return nil
}