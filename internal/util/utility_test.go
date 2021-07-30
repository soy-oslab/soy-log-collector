package util

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestRangeMapping(t *testing.T) {
	log.Println(RangeMapping(0, 10))
	log.Println(RangeMapping(1, 10))
	log.Println(RangeMapping(2, 10))
	log.Println(RangeMapping(3, 10))
	log.Println(RangeMapping(4, 10))
	log.Println(RangeMapping(5, 10))
	log.Println(RangeMapping(6, 10))
	log.Println(RangeMapping(7, 10))
	log.Println(RangeMapping(8, 10))
	log.Println(RangeMapping(9, 10))
}

func TestTimeSlice(t *testing.T) {
	ts := TimeSlice(1)
	fmt.Println(ts)

	ts = TimeSlice(time.Now().UnixNano())
	fmt.Println(ts)

	ts = TimeSlice(-1)
	fmt.Println(ts)
}
