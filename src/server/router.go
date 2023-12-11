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
	base.Get("/book/:id", f.GetOneBooks)
	base.Put("/book/:id", f.EditBook)
	base.Delete("/book/:id", f.DeleteBooks)
	base.Post("/book", f.CreateBooks)

}

func (f FiberServer) GetAllBooks(ctx *fiber.Ctx) error {
	//var books []services.Book //books me s has empty ระบุไทป์ เป็น book array
	shop, err := f.bookPg.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	return ctx.Status(200).JSON(shop)
}
func (f FiberServer) EditBooks(ctx *fiber.Ctx) error {
	var books []services.Book //books me s has empty ระบุไทป์ เป็น book array
	shop, err := f.bookPg.GetAll()
	if err != nil {
		log.Fatal(shop)
	}
	return ctx.Status(200).JSON(books)
}
func (f FiberServer) DeleteBooks(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := f.bookPg.Delete(id)
	if err != nil {
		log.Fatal(err)
	}
	//books me s has empty ระบุไทป์ เป็น book array

	return ctx.Status(204).JSON(fiber.Map{})
}

func (f FiberServer) CreateBooks(ctx *fiber.Ctx) error {
	CB := new(createbook)
	if err := ctx.BodyParser(CB); err != nil {
		return err
	}
	newBook := &services.Book{
		Name:     CB.BookName,
		Price:    CB.BookPrice,
		Category: CB.BookCategory,
	}
	if err := f.bookPg.Post(newBook); err != nil {
		log.Fatal(err)
	}
	return ctx.Status(201).JSON(CB)
}
func (f FiberServer) EditBook(ctx *fiber.Ctx) error {

	CEB := new(createbook)
	if err := ctx.BodyParser(CEB); err != nil {
		return err
	}
	bookid := ctx.Params("id")

	EditfindBook := &services.Book{
		Name:     CEB.BookName,
		Price:    CEB.BookPrice,
		Category: CEB.BookCategory,
	}
	if err := f.bookPg.Put(bookid, EditfindBook); err != nil {
		log.Fatal(err)

	}

	return ctx.Status(200).JSON(CEB)
}
func (f FiberServer) GetOneBooks(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	book, err := f.bookPg.GetOne(id)
	if err != nil {
		//log.Fatal(err)
		return ctx.Status(404).JSON(fiber.Map{"error": "Error 404 "})
	}
	return ctx.Status(200).JSON(book)
}
