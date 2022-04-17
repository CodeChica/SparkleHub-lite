package db

import (
	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/codechica/SparkleHub-lite/pkg/pls"
)

type Repository struct {
	sparkles []*domain.Sparkle
}

func NewRepository() *Repository {
	return &Repository{
		sparkles: []*domain.Sparkle{},
	}
}

func (s *Repository) All() []*domain.Sparkle {
	readonly := make([]*domain.Sparkle, len(s.sparkles))
	copy(readonly, s.sparkles)
	return readonly
}

func (s *Repository) Save(item *domain.Sparkle) error {
	if err := item.Validate(); err != nil {
		return err
	}

	if item.ID == "" {
		item.ID = pls.GenerateULID()
	}
	s.sparkles = append(s.sparkles, item)
	return nil
}
