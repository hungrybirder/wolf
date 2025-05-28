package fastpercentile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getNForPercentile(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(int64(57), getNForPercentile(99))
	assert.Equal(int64(113), getNForPercentile(99.5))
	assert.Equal(int64(562), getNForPercentile(99.9))
}
