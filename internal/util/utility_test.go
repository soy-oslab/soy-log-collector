package util

import (
	"fmt"
	"testing"
	"time"
)

func TestRangeMapping(t *testing.T) {
	fmt.Println(RangeMapping(0))
	fmt.Println(RangeMapping(1))
	fmt.Println(RangeMapping(2))
	fmt.Println(RangeMapping(3))
	fmt.Println(RangeMapping(4))
	fmt.Println(RangeMapping(5))
	fmt.Println(RangeMapping(6))
	fmt.Println(RangeMapping(7))
	fmt.Println(RangeMapping(8))
	fmt.Println(RangeMapping(9))
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

}
