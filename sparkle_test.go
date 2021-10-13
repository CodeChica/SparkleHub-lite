package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSparkle(t *testing.T) {
	t.Run("NewSparkle", func(t *testing.T) {
		t.Run("with a valid body", func(t *testing.T) {
			sparkle, err := NewSparkle("@monalisa for helping me with my homework!")

			assert.Nil(t, err)
			if err != nil {
				assert.Equal(t, "@monalisa", sparkle.Sparklee)
				assert.Equal(t, "for helping me with my homework!", sparkle.Reason)
			}
		})

		t.Run("with an empty body", func(t *testing.T) {
			sparkle, err := NewSparkle("")

			assert.Nil(t, sparkle)
			assert.NotNil(t, err)
			if err != nil {
				assert.Equal(t, "sparkle is empty", err.Error())
			}
		})

		t.Run("without a reason", func(t *testing.T) {
			sparkle, err := NewSparkle("@monalisa")

			assert.Nil(t, sparkle)
			assert.NotNil(t, err)
			if err != nil {
				assert.Equal(t, "sparkle is invalid", err.Error())
			}
		})

		t.Run("without a username", func(t *testing.T) {
			sparkle, err := NewSparkle("for helping me with my homework")

			assert.Nil(t, sparkle)
			assert.NotNil(t, err)
			if err != nil {
				assert.Equal(t, "sparkle is invalid", err.Error())
			}
		})
	})
}
