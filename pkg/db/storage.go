package db

import "github.com/codechica/SparkleHub-lite/pkg/domain"

type Storage struct {
	Sparkles []*domain.Sparkle
}

func NewStorage() *Storage {
	return &Storage{
		Sparkles: []*domain.Sparkle{},
	}
}

func (s *Storage) Save(item *domain.Sparkle) {
	s.Sparkles = append(s.Sparkles, item)
}
