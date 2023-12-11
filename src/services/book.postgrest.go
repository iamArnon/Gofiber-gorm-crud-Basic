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

func (g GormPg) GetAll() ([]Book, error) {
	var books []Book //books me s has empty ระบุไทป์ เป็น book array
	listBooks := g.database.Find(&books)
	if listBooks.Error != nil {
		log.Fatal(listBooks.Error)
	}
	return books, nil
}
func (g GormPg) GetOne(bookid string) (Book, error) {
	var book Book //books me s has empty ระบุไทป์ เป็น book array
	if err := g.database.First(&book, bookid).Error; err != nil {
		return Book{}, err
	}
	return book, nil
}

func (g GormPg) Delete(teenid string) error {

	var DelateBook Book
	if err := g.database.First(&DelateBook, teenid).Error; err != nil {
		return err
	}
	if err := g.database.Delete(&DelateBook).Error; err != nil {
		return err
	}

	return nil
}

func (g GormPg) Post(newBook *Book) error {
	if err := g.database.Create(newBook).Error; err != nil {
		return err
	}
	return nil
}

func (g GormPg) Put(bookid string, EditfindBook *Book) error {
	var EBook Book
	if err := g.database.First(&EBook, bookid).Error; err != nil {
		return err
	}
	EBook.Name = EditfindBook.Name
	EBook.Price = EditfindBook.Price
	EBook.Category = EditfindBook.Category

	if err := g.database.Save(&EBook).Error; err != nil {
		return err
	}
	return nil
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
