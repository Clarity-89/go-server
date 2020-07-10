package models

import "time"

type DbTodo struct {
	id        int64
	Title     string
	Content   string
	DueDate   int64
	CreatedAt int64
}

type TodoDTO struct {
	Title   string
	Content string
	DueDate time.Time
}
