package main

import (
	"github.com/pastel-lilac/clasick-api/internal/core/domain/model"
	"github.com/pastel-lilac/clasick-api/pkg/db"
	"log"
)

var (
	gormDriver db.IGormDriver
)

func init() {
	gormDriver = db.NewGormDriver()
	log.Printf("info: opend db connection \n")
	if err := gormDriver.Migrate(&model.Genre{},&model.Artist{},&model.Album{},&model.Music{}); err != nil {
		log.Printf("error: failed migrating database -> %v\n", err)
	}
}

func main() {

}
