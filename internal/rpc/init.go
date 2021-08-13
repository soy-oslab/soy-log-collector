package rpc

import (
	"os"
	"strconv"

	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// InitFlag is flag for MapTable is set
var InitFlag int

// ColdRing ring container for ColdPort rpc procedure
var ColdRing *ring.Ring

// ColdRingSize is default cold ring size
// Initialize with 10
var ColdRingSize int

func init() {
	ColdRingSize, _ = strconv.Atoi(os.Getenv("COLDPORTSIZE"))

	ColdRing = ring.New(ColdRingSize)
}
