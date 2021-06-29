package ring

import (
	"errors"

	"github.com/sheerun/queue"
)

// Ring container with fixed capacity
type Ring struct {
	ring     *queue.Queue
	capacity int
}

// DefaultSize is default Ring size
// Is set with 10
var DefaultSize int

func init() {
	DefaultSize = 10
}

// New returns Ring conatiner with capacity
// Default capacity is 10
func New(capacity ...int) *Ring {
	size := DefaultSize
	if len(capacity) > 0 {
		size = capacity[0]
	}
	ring := Ring{ring: queue.New(), capacity: size}
	return &ring
}

// Push returns error if elements exceed capacity
// Push args into Ring
func (t *Ring) Push(args interface{}) error {
	if t.capacity <= t.ring.Length() {
		return errors.New("Ring capacity is full!")
	}
	t.ring.Append(args)
	return nil
}

// Pop returns interfaces, error
// If there are no elements in Ring, return error
func (t *Ring) Pop() (interface{}, error) {
	if t.ring.Length() == 0 {
		return 0, errors.New("Ring is empty!")
	}
	return t.ring.Pop(), nil
}
