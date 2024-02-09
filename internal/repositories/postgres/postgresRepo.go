package postgres

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/schoolboybru/location-service/internal/domain"
)

type LocationRepository interface {
	AddLocation(location *domain.Location) error
	GetLocation(id string) (*domain.Location, error)
	GetLocationsByCity(id string) (*[]domain.Location, error)
	GetLocationsByCountry(id string) (*[]domain.Location, error)
}

type CommentRepository interface {
    AddComment(comment *domain.Comment) error
    GetComments(postId string) (*[]domain.Comment, error)
}

type DatabaseRepository interface {
    LocationRepository
    CommentRepository
}

type PostgresRepository struct {
	db *sqlx.DB
}

func New() (DatabaseRepository, error) {

    godotenv.Load(".env")

	dbhost := os.Getenv("HOST")
	dbuser := os.Getenv("USER")
	dbname := os.Getenv("NAME")
	dbPassword := os.Getenv("PASSWORD")
	dbport, err := strconv.Atoi(os.Getenv("PORT"))

    if err != nil {
        println(err)
    }

	conn := ConnString(dbhost, dbport, dbuser, dbPassword, dbname)
	database, err := sqlx.Connect("postgres", conn)

	if err != nil {
		return nil, err
	}

	err = database.Ping()

	if err != nil {
		return nil, err
	}

	repo := &PostgresRepository{db: database}

	return repo, nil

}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}
