package domain

type Service interface {
	AddLocation(location *Location) error
	GetLocation(id string) (*Location, error)
	GetLocationsByCity(id string) (*[]Location, error)
	GetLocationsByCountry(id string) (*[]Location, error)
}

type Repository interface {
	AddLocation(location *Location) error
	GetLocation(id string) (*Location, error)
	GetLocationsByCity(id string) (*[]Location, error)
	GetLocationsByCountry(id string) (*[]Location, error)
}

type service struct {
	r Repository
}

func New(repository Repository) Service {
	return &service{r: repository}
}

func (s *service) AddLocation(l *Location) error {

	return s.r.AddLocation(l)
}

func (s *service) GetLocation(id string) (*Location, error) {
	return s.r.GetLocation(id)
}

func (s *service) GetLocationsByCity(id string) (*[]Location, error) {
	return s.r.GetLocationsByCity(id)
}

func (s *service) GetLocationsByCountry(id string) (*[]Location, error) {
	return s.r.GetLocationsByCountry(id)
}
