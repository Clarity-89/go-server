package models

import "time"

type DbTodo struct {
	Id        int64
	Title     string
	Content   string
	DueDate   time.Time
	CreatedAt time.Time
}

type TodoDTO struct {
	Title   string
	Content string
	DueDate time.Time
}
