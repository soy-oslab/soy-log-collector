package util

import (
	"math"

	"github.com/soyoslab/soy_log_collector/internal/global"
)

// RangeMapping remap value into 255 using DefaultRingSize
func RangeMapping(arg int) uint8 {
	return uint8(math.Round(float64(arg) * 255 / float64(global.DefaultRingSize)))
}
