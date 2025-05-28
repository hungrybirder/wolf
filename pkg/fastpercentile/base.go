package fastpercentile

import "time"

type baseCounter struct {
	NForPercentile    int64 `json:"n_for_percentile,omitempty"`
	TotalCount        int64 `json:"total_count,omitempty"`
	LatencyMaxNS      int64 `json:"latency_max_ns,omitempty"`
	LatencySumOfMaxNS int64 `json:"latency_sum_of_max_ns,omitempty"`
}

func (b *baseCounter) Add(latency time.Duration) {
	latencyNS := latency.Nanoseconds()
	if latencyNS > b.LatencyMaxNS {
		b.LatencyMaxNS = latencyNS
	}
	b.TotalCount++
	if b.TotalCount%b.NForPercentile == 0 {
		b.LatencySumOfMaxNS += b.LatencyMaxNS
		b.LatencyMaxNS = 0
	}
}

func (b *baseCounter) Compute() time.Duration {
	if b.TotalCount < b.NForPercentile {
		return time.Duration(0)
	}
	return time.Duration(b.LatencySumOfMaxNS / (b.TotalCount / b.NForPercentile))
}

func (b *baseCounter) N() int64 {
	return b.NForPercentile
}
