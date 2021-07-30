package util

import (
	"math"
	"time"
)

// RangeMapping remap value into 255 using DefaultRingSize
func RangeMapping(arg int, ringSize int) uint8 {
	return uint8(math.Round(float64(arg) * 255 / float64(ringSize)))
}

// TimeSlice return Date, Time, NanoSecond in strings
// arg must unix nano time(int64) like time.Now().UnixNano()
func TimeSlice(arg int64) string {
	ts := time.Unix(0, arg).Format(time.RFC3339)
	return ts
}
