package global

import (
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// HotRing ring container for HotPort rpc procedure
var HotRing *ring.Ring

// ColdRing ring container for ColdPort rpc procedure
var ColdRing *ring.Ring

func init() {
	HotRing = ring.New()
	ColdRing = ring.New()
}
