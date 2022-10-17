package storage

import "bookshelf/models"

type Inmemory struct {
	Db *DB
}

type DB struct {
	InMemoryUserData []models.User
	InMemoryBookData []models.Book
}
