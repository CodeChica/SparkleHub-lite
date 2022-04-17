package db

import (
	"testing"

	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	storage := NewStorage()

	t.Run("Save", func(t *testing.T) {
		t.Run("an invalid Sparkle", func(t *testing.T) {
			err := storage.Save(&domain.Sparkle{Reason: "because"})

			assert.NotNil(t, err)
			assert.Equal(t, 0, len(storage.Sparkles))
		})

		t.Run("a valid Sparkle", func(t *testing.T) {
			err := storage.Save(&domain.Sparkle{Sparklee: "@monalisa", Reason: "because"})

			assert.Nil(t, err)
			assert.Equal(t, 1, len(storage.Sparkles))
			assert.Equal(t, "@monalisa", storage.Sparkles[0].Sparklee)
			assert.Equal(t, "because", storage.Sparkles[0].Reason)
		})
	})
}
