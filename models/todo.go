package models

type Todo struct {
	id int64
	Text string
	DueDate int64
	Category Category
}

type Category struct {
	name string
}
