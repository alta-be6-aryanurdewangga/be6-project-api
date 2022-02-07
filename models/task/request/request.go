package request

import "part3/models/book"

type BookRequest struct {
	Name      string `json:"name"`
	Publisher string `json:"publisher"`
	Author    string `json:"author"`
}

func (b *BookRequest) ToBook() book.Book {
	return book.Book{
		Name: b.Name,
		Publisher: b.Publisher,
		Author: b.Author,
	}
}