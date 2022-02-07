package task

import (
	"part3/models/task/response"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"createdAt" `
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`

	Task_ID	 uint	`json :"task_ID" gorm:"autoIncrement"`
	Name     string `json:"name" gorm:"not null;type:varchar(100)"`
	Priority int    `json:"priority" gorm:"not null;index"`
}

func (b *Book) ToBookResponse() response.BookResponse {
	return response.BookResponse{
		ID:        b.ID,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,

		Name:      b.Name,
		Publisher: b.Publisher,
		Author:    b.Author,
	}
}
