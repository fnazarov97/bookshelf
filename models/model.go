package models

import "time"

type User struct {
	ID        string     `json:"id" binding:"required"`
	Name      string     `json:"name" binding:"required"`
	Key       string     `json:"key" binding:"required"`
	Secret    string     `json:"secret" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}
type CreateUserModel struct {
	Name   string `json:"name" binding:"required"`
	Key    string `json:"key" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}
type GetUserModel struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" binding:"required"`
	Key       string    `json:"key" binding:"required"`
	Secret    string    `json:"secret" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type Book struct {
	ID        string     `json:"id" binding:"required"`
	Isbn      string     `json:"isbn" binding:"required"`
	Title     string     `json:"title" binding:"required"`
	Author    string     `json:"author" binding:"required"`
	Published uint8      `json:"published" binding:"required"`
	Pages     uint8      `json:"pages" binding:"required"`
	Status    uint8      `json:"status" binding:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type CreateBookModel struct {
	Isbn      string `json:"isbn" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Published uint8  `json:"published" binding:"required"`
	Pages     uint8  `json:"pages" binding:"required"`
	Status    uint8  `json:"status" binding:"required"`
}
type GetBookModel struct {
	ID        string `json:"id" binding:"required"`
	Isbn      string `json:"isbn" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Author    string `json:"author" binding:"required"`
	Published uint8  `json:"published" binding:"required"`
	Pages     uint8  `json:"pages" binding:"required"`
	Status    uint8  `json:"status" binding:"required"`
}

type BookI struct {
	Isbn string
}
