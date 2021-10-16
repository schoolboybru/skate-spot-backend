package postgres

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/schoolboybru/location-service/internal/repositories"

	"github.com/schoolboybru/location-service/internal/domain"
)

type PostgresRepository struct {
	db *sqlx.DB
}

func New() (repositories.LocationRepository, error) {

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

	repo := &PostgresRepository{db: database}

	return repo, nil

}

func ConnString(host string, port int, user string, password string, dbName string) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}

func (p *PostgresRepository) AddLocation(l *domain.Location) error {

	client := p.db

	defer client.Close()

	_, err := client.NamedExec(`INSERT INTO LOCATION (id, name, longitude, latitude, date, city, country) VALUES (:id, :name, :longitude, :latitude, :date, :city, :country)`, &l)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) GetLocation(id string) (*domain.Location, error) {

	var location domain.Location

	client := p.db

	defer client.Close()

	err := client.Get(&location, "SELECT * FROM Location WHERE id=$1", id)

	if err != nil {
		println(err.Error())
	}

	return &location, nil
}

func (p *PostgresRepository) GetLocationsByCity(id string) (*[]domain.Location, error) {

	var locations []domain.Location

	client := p.db

	defer client.Close()

	err := client.Select(&locations, `SELECT * FROM LOCATION WHERE city=$1`, id)

	if err != nil {
		return nil, err
	}

	return &locations, nil
}

func (p *PostgresRepository) GetLocationsByCountry(id string) (*[]domain.Location, error) {

	var locations []domain.Location

	client := p.db

	defer client.Close()

	err := client.Select(&locations, `SELECT * FROM LOCATION WHERE country=$1`, id)

	if err != nil {
		return nil, err
	}

	return &locations, nil
}
