package services

import (
	"database/sql"
	"fmt"
	"go-server/models"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "todo"
	password = "password"
	dbname   = "go"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	// TODO Better error handling
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title varchar(45), content varchar(450) NOT NULL, due_date timestamp, created_at timestamp)")
	if err != nil {
		log.Print("Could not create todos table", err)
	}

	return db
}

func SaveTodo(todo models.Todo) error {
	db := Init()
	_, err := db.Exec("INSERT INTO todos (todo) VALUES($1)", todo)
	return err
}
