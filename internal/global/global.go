package global

import (
	"context"

	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
	"github.com/soyoslab/soy_log_collector/pkg/server"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// DefaultRingSize is default ring buffer size
// Initialize with 10
var DefaultRingSize int

// HotRing ring container for HotPort rpc procedure
var HotRing *ring.Ring

// ColdRing ring container for ColdPort rpc procedure
var ColdRing *ring.Ring

// RedisServer is used for rpc data server with Hot/Cold port.
var RedisServer server.Server

// Compressor is for rpc.
var Compressor compressor.Compressor

var ctx context.Context

func init() {
	DefaultRingSize = 10
	ctx = context.Background()
	Compressor = &(compressor.GzipComp{})
	HotRing = ring.New(DefaultRingSize)
	ColdRing = ring.New(DefaultRingSize)
	RedisServer = server.New(ctx)
}
