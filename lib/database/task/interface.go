package task

import "part3/models/task"

type Task interface {
	Create(newTask task.Task) (task.Task, error)
	// GetById(id int) (book.Book, error)
	// UpdateById(id int, bookReg request.BookRequest) (book.Book, error)
	// DeleteById(id int) (gorm.DeletedAt, error)
}
