package main

import (
	"bookshelf/handlers"
	"bookshelf/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	im := storage.Inmemory{
		Db: &storage.DB{},
	}
	h := handlers.Handler{
		IM: im,
	}
	router := gin.Default()
	router.POST("/signup", h.CreateUser)
	router.GET("/myself", h.GetUser)
	router.POST("/books", h.CreateBook)
	router.GET("/books", h.GetAllBooks)

	router.Run(":5000")

}
