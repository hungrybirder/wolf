package fastpercentile

import (
	"errors"
	"time"
)

type Percentile float64

const (
	P99  Percentile = 99.0
	P995 Percentile = 99.5
	P999 Percentile = 99.9
)

var ErrInvalidPercentile = errors.New("invalid percentile")

// LatencyPercentileCounter Interface
// Provide a way to add latency and compute percentile latency
type LatencyPercentileCounter interface {
	// Add one lantency
	Add(latency time.Duration)

	// Compute
	// If the total count is less than the nForPercentile, return 0
	// nForPercentile is 57 for P99
	// nForPercentile is 113 for P995
	// nForPercentile is 562 for P999
	Compute() time.Duration
}

type FastPercentile struct {
	baseCounter
}

// New create a new FastPercentile
func New(p Percentile) (*FastPercentile, error) {
	if p <= 0 || p >= 100 {
		return nil, ErrInvalidPercentile
	}
	fp := &FastPercentile{
		baseCounter: baseCounter{
			NForPercentile: getNForPercentile(float64(p)),
		},
	}
	return fp, nil
}

// NewP99 create a new FastPercentile with P99
func NewP99() *FastPercentile {
	fp, _ := New(P99)
	return fp
}

// NewP995 create a new FastPercentile with P995
func NewP995() *FastPercentile {
	fp, _ := New(P995)
	return fp
}

// NewP999 create a new FastPercentile with P999
func NewP999() *FastPercentile {
	fp, _ := New(P999)
	return fp
}
