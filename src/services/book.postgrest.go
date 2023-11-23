package services

import (
	"log"

	"gorm.io/gorm"
)

type GormPg struct {
	database *gorm.DB
}

type Book struct {
	gorm.Model

	Name     string `gorm:"not null"`
	Price    int    `gorm:"not null"`
	Category string `gorm:"not null"`
}

func (g GormPg) GetAll() (*gorm.DB, error) {
	var books []Book //books me s has empty ระบุไทป์ เป็น book array
	listBooks := g.database.Find(&books)
	if listBooks.Error != nil {
		log.Fatal(listBooks)
	}
	return listBooks, nil
}

func SetUpPosgresql(db *gorm.DB) GormPg {
	g := GormPg{
		database: db,
	}
	if err := g.database.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Error migrate")
	} // create table
	log.Println("Connected to the database")
	return g
}
