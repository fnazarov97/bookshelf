package handlers

import (
	"bookshelf/models"
	"fmt"
	"net/http"

	goisbn "github.com/abx123/go-isbn"
	"github.com/gin-gonic/gin"
)

func (h Handler) CreateBook(c *gin.Context) {
	_, check := h.Check(c)
	if check {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{Error: "Unauthorized"})
		return
	}
	gi := goisbn.NewGoISBN(goisbn.DEFAULT_PROVIDERS)
	var B models.BookI

	if err := c.ShouldBindJSON(&B); err != nil {
		fmt.Println(B.Isbn, err)
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "invalid input"})
		return
	}
	fmt.Println(B.Isbn)
	books, err := h.IM.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	for _, b := range books {
		if b.Isbn == B.Isbn {
			c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "isbn of book already exist"})
			return
		}
	}
	book, err := gi.Get(B.Isbn)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	b, err := h.IM.CreateBook(book)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "error created book"})
		return
	}
	c.JSON(http.StatusOK, models.JSONBookResponse{
		Data: models.Data{
			Book:   b,
			Status: 0,
		},
		IsOk:    true,
		Message: "ok",
	})
}

func (h Handler) GetAllBooks(c *gin.Context) {
	_, check := h.Check(c)
	if check {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{Error: "Unauthorized"})
		return
	}
	Books, err := h.IM.GetBooks()
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: "error get books"})
		return
	}
	c.JSON(http.StatusOK, models.JSONGetBooksResponse{
		Data:    Books,
		IsOk:    true,
		Message: "ok",
	})

}

func (h Handler) UpdateBook(c *gin.Context) {
	_, check := h.Check(c)
	if check {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{Error: "Unauthorized"})
		return
	}
	var body = models.CreateBookModel{}
	var id = c.Param("id")
	book, err := h.IM.UpdateBook(id, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, models.JSONUpdateBookResponse{
		Data:    book,
		IsOk:    true,
		Message: "ok",
	})

}

func (h Handler) DeleteBook(c *gin.Context) {
	_, check := h.Check(c)
	if check {
		c.JSON(http.StatusUnauthorized, models.JSONErrorResponse{Error: "Unauthorized"})
		return
	}
	idStr := c.Param("id")

	err := h.IM.DeleteBook(idStr)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": err})
		return
	}
	c.JSON(http.StatusOK, models.JSONDeleteBookResponse{
		Data:    "Saccesfully deleted",
		IsOk:    true,
		Message: "ok",
	})
}
