package postgres

import "github.com/schoolboybru/skate-spot/internal/domain"

type PostRespository interface {
	AddPost(post *domain.Post) error
	GetPost(id string) (*domain.Post, error)
}

func (p *PostgresRepository) AddPost(post *domain.Post) error {
	client := p.db

	_, err := client.NamedExec(`INSERT INTO LOCATION (id, user_id, location_id, created_at, likes) VALUES (:id, :user_id, :location_id, :created_at, :likes)`, &post)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) GetPost(id string) (*domain.Post, error) {
	var post domain.Post

	client := p.db

	err := client.Get(&post, "SELECT * FROM post WHERE id=$1", id)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
