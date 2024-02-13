package postgres

import "github.com/schoolboybru/skate-spot/internal/domain"

type CommentRepository interface {
	AddComment(comment *domain.Comment) error
	GetComments(postId string) (*[]domain.Comment, error)
}

func (p *PostgresRepository) AddComment(c *domain.Comment) error {

	client := p.db

	_, err := client.NamedExec(`INSERT INTO COMMENT (id, post_id, user_id, value, created_at) VALUES (:id, :post_id, :user_id, :value, :created_at)`, &c)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostgresRepository) GetComments(postId string) (*[]domain.Comment, error) {
	var comments []domain.Comment

	client := p.db

	err := client.Select(&comments, `SELECT * FROM COMMENT WHERE post_id = $1`, postId)

	if err != nil {
		return nil, err
	}

	return &comments, nil
}
