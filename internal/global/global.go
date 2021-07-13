package global

import (
	"context"
	"flag"
	"os"

	"github.com/smallnest/rpcx/client"
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
var RedisServer *server.Server

// Compressor is for rpc.
var Compressor compressor.Compressor

// SoyLogExplorer rpc server client
var SoyLogExplorer client.XClient

var exploreraddr = os.Getenv("EXPLORERSERVER")

var addr = flag.String("addr", exploreraddr, "server address")

var ctx context.Context

func init() {
	DefaultRingSize = 10
	ctx = context.Background()
	Compressor = &(compressor.GzipComp{})
	HotRing = ring.New(DefaultRingSize)
	ColdRing = ring.New(DefaultRingSize)
	RedisServer = server.New(ctx)
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	SoyLogExplorer = client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
}
