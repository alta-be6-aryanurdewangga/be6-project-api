package book

import (
	"part3/models/book"
	"part3/models/book/request"

	"gorm.io/gorm"
)

type Book interface {
	Create(newBook book.Book) (book.Book, error)
	GetById(id int) (book.Book, error)
	UpdateById(id int, bookReg request.BookRequest) (book.Book, error)
	DeleteById(id int) (gorm.DeletedAt, error)
}
