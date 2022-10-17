package storage

import (
	"bookshelf/models"
	"errors"
	"time"

	goisbn "github.com/abx123/go-isbn"
)

func (im Inmemory) CreateBook(book *goisbn.Book) (models.Book, error) {
	var fullBook models.Book
	if book.Title == "" {
		return fullBook, errors.New("book empty")
	}
	fullBook.Isbn = book.IndustryIdentifiers.ISBN13
	fullBook.Title = book.Title
	fullBook.Author = book.Authors[0]
	fullBook.Pages = uint8(book.PageCount)
	fullBook.CreatedAt = time.Now()
	im.Db.InMemoryBookData = append(im.Db.InMemoryBookData, fullBook)
	return fullBook, nil
}

func (im Inmemory) GetBooks() (Books []models.GetBookModel, err error) {
	var getBook models.GetBookModel

	for _, book := range im.Db.InMemoryBookData {
		getBook.ID = book.ID
		getBook.Isbn = book.Isbn
		getBook.Title = book.Title
		getBook.Author = book.Author
		getBook.Published = book.Published
		getBook.Pages = book.Pages
		Books = append(Books, getBook)
	}
	return

}

func (im Inmemory) UpdateBook(id string, body models.CreateBookModel) (Book models.GetBookModel, err error) {
	for _, book := range im.Db.InMemoryBookData {
		if book.ID == id {
			Book.ID = book.ID
			Book.Author = book.Author
			Book.Isbn = book.Isbn
			Book.Published = book.Published
			Book.Pages = book.Pages
			Book.Title = book.Title
			Book.Status = body.Status
			return Book, nil
		}
	}
	return Book, errors.New("book not found")

}

func (im Inmemory) DeleteBook(id string) error {
	for i, v := range im.Db.InMemoryBookData {
		if v.ID == id {
			if v.DeletedAt != nil {
				return errors.New("book alredy deleted")
			}
			t := time.Now()
			v.DeletedAt = &t
			im.Db.InMemoryBookData[i] = v
			return nil
		}
	}
	return errors.New("not found for delete")
}
