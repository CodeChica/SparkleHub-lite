package db

import (
	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/codechica/SparkleHub-lite/pkg/pls"
)

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

	if item.ID == "" {
		item.ID = pls.GenerateULID()
	}
	s.Sparkles = append(s.Sparkles, item)
	return nil
}
