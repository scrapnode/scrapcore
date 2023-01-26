package utils_test

import (
	"github.com/scrapnode/scrapcore/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeRoundUp(t *testing.T) {
	t.Run("2022-02-22T22:22:22.000Z -> 2022-02-22T23:00:00.000Z", func(t *testing.T) {
		start := time.UnixMilli(1645568542000)
		end := utils.TimeRoundUp(start, time.Hour)
		assert.Equal(t, int64(1645570800000), end.UnixMilli())
	})
	t.Run("2022-02-22T22:32:22.000Z -> 2022-02-22T23:00:00.000Z", func(t *testing.T) {
		start := time.UnixMilli(1645569142000)
		end := utils.TimeRoundUp(start, time.Hour)
		assert.Equal(t, int64(1645570800000), end.UnixMilli())
	})
	t.Run("2022-02-22T23:00:00.000Z -> 2022-02-22T23:00:00.000Z", func(t *testing.T) {
		start := time.UnixMilli(1645570800000)
		end := utils.TimeRoundUp(start, time.Hour)
		assert.Equal(t, int64(1645570800000), end.UnixMilli())
	})
}
