package handlers

import (
	"bookshelf/models"

	goisbn "github.com/abx123/go-isbn"
)

type Handler struct {
	IM StorageI
}

type StorageI interface {
	AddUser(id string, entity models.CreateUserModel) error
	GetUser(id string) (models.GetUserModel, error)
	CreateBook(book *goisbn.Book) (models.Book, error)
	GetBooks() (Books []models.GetBookModel, err error)
	UpdateBook(id string, body models.CreateBookModel) (Book models.GetBookModel, err error)
	DeleteBook(id string) error
}
