package storage

import (
	"database/sql"
	"fmt"
	"go-server/models"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "todo"
	password = "password"
	dbname   = "go"
)

func Init() Storage {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	// TODO Better error handling
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title varchar(45), content varchar(450) NOT NULL, due_date timestamp with time zone, created_at timestamp with time zone)")
	if err != nil {
		log.Print("Could not create todos table", err)
	}

	return Storage{db: db}
}

func (s Storage) SaveTodo(todo models.TodoDTO) error {
	query := "INSERT INTO todos (title, content, due_date, created_at) VALUES ($1, $2, $3, $4)"
	_, err := s.db.Exec(query, todo.Title, todo.Content, todo.DueDate, time.Now())
	return err
}
