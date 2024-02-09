package postgres

import "github.com/schoolboybru/location-service/internal/domain"

func (p *PostgresRepository) AddLocation(l *domain.Location) error {

	client := p.db

	_, err := client.NamedExec(`INSERT INTO LOCATION (id, name, longitude, latitude, date, city, country) VALUES (:id, :name, :longitude, :latitude, :date, :city, :country)`, &l)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) GetLocation(id string) (*domain.Location, error) {

	var location domain.Location

	client := p.db

	err := client.Get(&location, "SELECT * FROM Location WHERE id=$1", id)

	if err != nil {
		println(err.Error())
	}

	return &location, nil
}

func (p *PostgresRepository) GetLocationsByCity(id string) (*[]domain.Location, error) {

	var locations []domain.Location

	client := p.db

	err := client.Select(&locations, `SELECT * FROM LOCATION WHERE city=$1`, id)

	if err != nil {
		return nil, err
	}

	return &locations, nil
}

func (p *PostgresRepository) GetLocationsByCountry(id string) (*[]domain.Location, error) {

	var locations []domain.Location

	client := p.db

	err := client.Select(&locations, `SELECT * FROM LOCATION WHERE country=$1`, id)

	if err != nil {
		return nil, err
	}

	return &locations, nil
}
