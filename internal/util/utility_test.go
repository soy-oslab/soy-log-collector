package util

import (
	"log"
	"testing"
	"time"
)

func TestRangeMapping(t *testing.T) {
	log.Println(RangeMapping(0))
	log.Println(RangeMapping(1))
	log.Println(RangeMapping(2))
	log.Println(RangeMapping(3))
	log.Println(RangeMapping(4))
	log.Println(RangeMapping(5))
	log.Println(RangeMapping(6))
	log.Println(RangeMapping(7))
	log.Println(RangeMapping(8))
	log.Println(RangeMapping(9))
}

func TestTimeSlice(t *testing.T) {
	_, _, _, err := TimeSlice(1)
	if err != nil {
		t.Error(err)
	}

	_, _, _, err = TimeSlice(time.Now().UnixNano())
	if err != nil {
		t.Error(err)
	}

	_, _, _, err = TimeSlice(-1)
	if err != nil {
		t.Error(err)
	}
}
