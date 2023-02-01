package utils_test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/scrapnode/scrapcore/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomString(t *testing.T) {
	t.Run("1 char", func(t *testing.T) {
		assert.Equal(t, 1, len(utils.RandomString(1)))
	})
	t.Run("multiple chars", func(t *testing.T) {
		count := gofakeit.IntRange(10, 100)
		assert.Equal(t, count, len(utils.RandomString(count)))
	})
}
