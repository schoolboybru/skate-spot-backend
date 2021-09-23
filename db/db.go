package db

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/schoolboybru/location-service/internal/adding"
)

type postgresRepository struct {
	db *sqlx.DB
}

func New() (adding.Repository, error) {

	godotenv.Load(".env")

	dbhost := os.Getenv("HOST")
	dbport, err := strconv.Atoi(os.Getenv("PORT"))
	dbuser := os.Getenv("PGUSER")
	dbname := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")

	conn := ConnString(dbhost, dbport, dbuser, dbPassword, dbname)
	database, err := sqlx.Connect("postgres", conn)

	if err != nil {
		return nil, err
	}

	err = database.Ping()

	if err != nil {
		return nil, err
	}

	repo := &postgresRepository{db: database}

	return repo, nil

}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}

func (p *postgresRepository) AddLocation(l *adding.Location) error {

	client := p.db

	defer client.Close()

	_, err := client.NamedExec(`INSERT INTO LOCATION (id, name, longitude, latitude, date, city, country) VALUES (:id, :name, :longitude, :latitude, :date, :city, :country)`, &l)

	if err != nil {
		return err
	}

	return nil
}
