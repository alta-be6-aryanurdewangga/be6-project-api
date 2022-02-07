package book

import (
	"part3/models/book/response"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Name     string `json:"name" gorm:"not null;type:varchar(100)"`
	Publisher    string `json:"publisher" gorm:"not null;type:varchar(100)"`
	Author string `json:"author" gorm:"not null;type:varchar(100)"`
}

func (b *Book) ToBookResponse() response.BookResponse {
	return response.BookResponse{
		ID: b.ID,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,

		Name: b.Name,
		Publisher: b.Publisher,
		Author: b.Author,
	}
}