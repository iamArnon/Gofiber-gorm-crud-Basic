package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type config struct {
	Port     string `env:"PORT" envDefault:"3000"`
	POSTGRES struct {
		Port     string `env:"POSTGRES_PORT" envDefault:"5000"`
		Host     string `env:"POSTGRES_HOST" envDefault:"localhost"`
		User     string `env:"POSTGRES_USER" envDefault:"backend"`
		Password string `env:"POSTGRES_PASSWORD" envDefault:"1234567890"`
		Database string `env:"POSTGRES_DATABASE" envDefault:"bookshop"`
	}
}
type book struct {
	gorm.Model
	Name     string `gorm:"not null"`
	Price    int    `gorm:"not null"`
	Category string `gorm:"not null"`
}
type createbook struct {
	BookName     string `json:"BookName"`
	BookPrice    int    `json:"BookPrice"`
	BookCategory string `json:"BookCategory"`
}

func main() {
	//load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// set config
	//ให้ตัวแปรกับenv match กัน
	var cfg config
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	//connect database
	dsn := "host=localhost user=backend password=1234567890 dbname=bookshop port=5000 sslmode=disable "
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	app := fiber.New()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	log.Println("Connected to the database")
	if err := db.AutoMigrate(&book{}); err != nil {
		log.Fatal("Error migrate")
	} // create table

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Post("/addbook", func(ctx *fiber.Ctx) error {
		CB := new(createbook)
		if err := ctx.BodyParser(CB); err != nil {
			return err
		}
		err := db.Model(&book{}).Create(&book{
			Name:     CB.BookName,
			Price:    CB.BookPrice,
			Category: CB.BookCategory,
		}).Error
		if err != nil {
			log.Fatal("Error naja", err)
		}
		return nil
	})
	app.Get("/shop", func(ctx *fiber.Ctx) error {
		var books []book //books me s has empty ระบุไทป์ เป็น book array
		db.Find(&books)

		return ctx.JSON(books)

	})
	app.Put("/editbook/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var EditBook book
		if err := db.First(&EditBook, id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString("Book not found")
		}
		if err := db.Model(&EditBook).Updates(book{Name: "hello", Price: 18, Category: "anime"}).Error; err != nil {
			return errS
		}
		CBE := new(createbook)
		if err := ctx.BodyParser(CBE); err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("invalid json")
		}
		EditBook.Name = CBE.BookName
		EditBook.Price = CBE.BookPrice
		EditBook.Category = CBE.BookCategory
		if err := db.Save(&EditBook).Error; err != nil { //ถ้าไม่เท่ากับ nil ก็จะไม่รีเทิร์น
			return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to update book")
		}
		return ctx.Status(fiber.StatusOK).SendString("Edit book suscess!")
	})
	app.Delete("delete/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var DelateBook book
		if err := db.First(&DelateBook, id).Error; err != nil {
			return ctx.Status(fiber.StatusNotFound).SendString("Book not found")
		}
		if err := db.Delete(&DelateBook).Error; err != nil {
			return ctx.Status(fiber.StatusBadRequest).SendString("Can't Delete data")
		}
		return ctx.Status(fiber.StatusOK).SendString("Book delete sucessfully")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatal("connect server suessfully")
	}
}
