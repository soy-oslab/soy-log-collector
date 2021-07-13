package util

import (
	"errors"
	"math"
	"strings"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
)

// RangeMapping remap value into 255 using DefaultRingSize
func RangeMapping(arg int) uint8 {
	return uint8(math.Round(float64(arg) * 255 / float64(global.DefaultRingSize)))
}

// TimeSlice return Date, Time, NanoSecond in strings
// arg must unix nano time(int64) like time.Now().UnixNano()
func TimeSlice(arg int64) (string, string, string, error) {
	t := time.Unix(0, arg).String()
	times := strings.Split(t, " ")

	if len(times) < 2 {
		err := errors.New("Time format is not proper")
		return "", "", "", err
	}

	date := times[0]
	now := times[1]
	times = strings.Split(now, ".")

	if len(times) < 2 {
		err := errors.New("Time format is not proper")
		return "", "", "", err
	}
	sec := times[0]
	nano := times[1]

	return date, sec, nano, nil
}
