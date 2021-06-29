package global

import (
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// DefaultRingSize is default ring buffer size
// Initialize with 10
var DefaultRingSize int

// HotRing ring container for HotPort rpc procedure
var HotRing *ring.Ring

// ColdRing ring container for ColdPort rpc procedure
var ColdRing *ring.Ring

func init() {
	DefaultRingSize = 10
	HotRing = ring.New(DefaultRingSize)
	ColdRing = ring.New(DefaultRingSize)
}
