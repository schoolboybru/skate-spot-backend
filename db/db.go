package db

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type Db struct {
	*sqlx.DB
}

func New() (*Db, error) {

	dbhost := os.Getenv("HOST")
	dbport, err := strconv.Atoi(os.Getenv("PORT"))
	//dbuser := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")

	conn := ConnString(dbhost, dbport, "root", dbPassword, dbname)
	db, err := sqlx.Connect("postgres", conn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return &Db{db}, nil

}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
