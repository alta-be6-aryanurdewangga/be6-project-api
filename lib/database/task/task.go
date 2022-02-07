package task

import (
	"part3/models/task"

	"gorm.io/gorm"
)

type TaskDb struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (td *TaskDb) Create(newTask task.Task) (task.Task, error) {
	if err := td.db.Create(&newTask).Error; err != nil {
		return newTask, err
	}
	return newTask, nil
}

// func (bd *BookDb) GetById(id int) (book.Book, error) {
// 	book := book.Book{}

// 	if err := bd.db.Model(&book).Where("id = ?", id).First(&book).Error; err != nil {
// 		return book, err
// 	}

// 	return book, nil
// }

// func (bd *BookDb) UpdateById(id int, bookReg request.BookRequest) (book.Book, error) {
// 	_, err := bd.GetById(id)

// 	if err != nil {
// 		return book.Book{}, err
// 	}

// 	bd.db.Model(book.Book{ID: uint(id)}).Updates(book.Book{Name: bookReg.Name, Publisher: bookReg.Publisher, Author: bookReg.Author})

// 	book := bookReg.ToBook()

// 	return book, nil
// }

// func (bd *BookDb) DeleteById(id int) (gorm.DeletedAt, error)  {
// 	book := book.Book{}
// 	_, err := bd.GetById(id)

// 	if err != nil{
// 		return book.DeletedAt, err
// 	}

// 	bd.db.Model(&book).Where("id = ?", id).Delete(&book)

// 	return book.DeletedAt, nil
// }

// func (bd *BookDb) GetAll() ([]response.BookResponse, error) {
// 	bookRespArr := []response.BookResponse{}

// 	if err := bd.db.Model(book.Book{}).Limit(5).Find(&bookRespArr).Error ; err != nil {
// 		return nil, err
// 	}

// 	return bookRespArr, nil
// }