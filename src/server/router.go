package server

import (
	"log"

	"Bookshop/src/services"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct {
	server *fiber.App
	config *ServerConfig
	bookPg services.GormPg
}

type ServerConfig struct {
	AppVersion    string
	ListenAddress string
}

type createbook struct {
	BookName     string `json:"BookName"`
	BookPrice    int    `json:"BookPrice"`
	BookCategory string `json:"BookCategory"`
}

func New(sc *ServerConfig, db services.GormPg) *FiberServer {
	app := fiber.New()
	f := &FiberServer{
		app,
		sc,
		db,
	}

	f.SetupRouteHttp(app)

	return f
}

func (f FiberServer) Start() {
	err := f.server.Listen(f.config.ListenAddress)
	if err != nil {
		log.Fatal("Server error")
	}
	log.Println("Fiber server is start up")
}

func (f FiberServer) SetupRouteHttp(base fiber.Router) {
	base.Get("/book", f.GetAllBooks)
	base.Put("/book/:id", func(ctx *fiber.Ctx) error {
		// id := ctx.Params("id")
		// var EditBook services.Book
		// if err := db.First(&EditBook, id).Error; err != nil {
		// 	return ctx.Status(fiber.StatusNotFound).SendString("Book not found")
		// }

		// CBE := new(createbook)
		// if err := ctx.BodyParser(CBE); err != nil {
		// 	return ctx.Status(fiber.StatusBadRequest).SendString("invalid json")
		// }
		// EditBook.Name = CBE.BookName
		// EditBook.Price = CBE.BookPrice
		// EditBook.Category = CBE.BookCategory
		// if err := db.Save(&EditBook).Error; err != nil { //ถ้าไม่เท่ากับ nil ก็จะไม่รีเทิร์น
		// 	return ctx.Status(fiber.StatusInternalServerError).SendString("Failed to update book")
		// }
		return ctx.Status(fiber.StatusOK).SendString("Edit book suscess!")
	})
	base.Delete("/book/:id", func(ctx *fiber.Ctx) error {
		// id := ctx.Params("id")
		// var DelateBook services.Book
		// if err := db.First(&DelateBook, id).Error; err != nil {
		// 	return ctx.Status(fiber.StatusNotFound).SendString("Book not found")
		// }
		// if err := db.Delete(&DelateBook).Error; err != nil {
		// 	return ctx.Status(fiber.StatusBadRequest).SendString("Can't Delete data")
		// }
		return ctx.Status(fiber.StatusOK).SendString("Book delete sucessfully")
	})
	base.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	base.Post("/book", func(ctx *fiber.Ctx) error {
		// CB := new(createbook)
		// if err := ctx.BodyParser(CB); err != nil {
		// 	return err
		// }
		// err := db.Model(&services.Book{}).Create(&services.Book{
		// 	Name:     CB.BookName,
		// 	Price:    CB.BookPrice,
		// 	Category: CB.BookCategory,
		// }).Error
		// if err != nil {
		// 	log.Fatal("Error naja", err)
		// }
		return nil
	})
}

func (f FiberServer) GetAllBooks(ctx *fiber.Ctx) error {
	var books []services.Book //books me s has empty ระบุไทป์ เป็น book array
	shop, err := f.bookPg.GetAll()
	if err != nil {
		log.Fatal(shop)
	}
	return ctx.JSON(books)
}
