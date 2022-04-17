package db

import "github.com/codechica/SparkleHub-lite/pkg/domain"

type Repository struct {
	Sparkles []*domain.Sparkle
}

func NewRepository() *Repository {
	return &Repository{
		Sparkles: []*domain.Sparkle{},
	}
}

func (s *Repository) Save(item *domain.Sparkle) error {
	if err := item.Validate(); err != nil {
		return err
	}

	s.Sparkles = append(s.Sparkles, item)
	return nil
}
