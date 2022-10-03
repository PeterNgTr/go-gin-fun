package main

import (
	"fmt"
	"github.com/peterngtr/go-gin-fun/controllers"
	"github.com/peterngtr/go-gin-fun/models"
	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/peterngtr/go-gin-fun/database"

	"github.com/gin-gonic/gin"
)
const DEFAULT_DB = "go-gin-fun.db"

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", DEFAULT_DB)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DBConn.AutoMigrate(models.Book{})
	faker := faker.New()

	for i := 1; i < 5; i++ {
		var book models.Book
		book.Title = faker.Person().Name()
		book.Author = faker.Person().Name()
		database.DBConn.Create(&book)
	}
}


func main() {
	r := gin.Default()

	initDatabase()

	// Routes
	r.GET("/books", controllers.FindBooks)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// Run the server
	r.Run()
}