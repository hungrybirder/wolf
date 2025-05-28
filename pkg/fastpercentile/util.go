package fastpercentile

import "math"

func getNForPercentile(percentile float64) int64 {
	var n int64 = 1
	var expectedValue float64 = 0
	percentileValue := -math.Log(1 - percentile/100)
	for expectedValue < percentileValue {
		expectedValue = expectedValue + 1.0/float64(n)
		n++
	}
	return n
}
