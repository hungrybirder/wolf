package fastpercentile

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert := assert.New(t)
	fp, err := New(P99)
	assert.NoError(err)
	assert.Equal(int64(57), fp.N())
	for i := range 57 * 3 {
		fp.Add(time.Duration(i+1) * time.Millisecond)
		if i < int(fp.N()-1) {
			assert.Equal(time.Duration(0), fp.Compute())
		} else {
			t.Logf("i: %d, fp.Compute(): %v", i, fp.Compute())
		}
	}
	assert.Equal(time.Duration(time.Millisecond*114), fp.Compute())
}
