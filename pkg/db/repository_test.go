package db

import (
	"testing"

	"github.com/codechica/SparkleHub-lite/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	storage := NewRepository()

	t.Run("Save", func(t *testing.T) {
		t.Run("an invalid Sparkle", func(t *testing.T) {
			err := storage.Save(&domain.Sparkle{Reason: "because"})

			assert.NotNil(t, err)
			assert.Equal(t, 0, len(storage.sparkles))
		})

		t.Run("a valid Sparkle", func(t *testing.T) {
			err := storage.Save(&domain.Sparkle{Sparklee: "@monalisa", Reason: "because"})

			assert.Nil(t, err)
			assert.Equal(t, 1, len(storage.sparkles))
			assert.NotEqual(t, "", storage.sparkles[0].ID)
			assert.Equal(t, "@monalisa", storage.sparkles[0].Sparklee)
			assert.Equal(t, "because", storage.sparkles[0].Reason)
		})
	})
}
