package main

import (
	"fmt"

	"github.com/ariskusnulwiditama/go-fiber-tutorial/book"
	"github.com/ariskusnulwiditama/go-fiber-tutorial/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Get("/api/v1/book", book.GetBooks)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")

	if err != nil {
		panic("error mas")
	}

	fmt.Println("sukses yeah")
	database.DBConn.AutoMigrate(&book.Book{})
}
