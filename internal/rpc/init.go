package rpc

import (
	"flag"
	"os"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// MapTable is used for log file indexes
var MapTable map[string][]string

// InitFlag is flag for MapTable is set
var InitFlag int

// HotRing ring container for HotPort rpc procedure
var HotRing *ring.Ring

// ColdRing ring container for ColdPort rpc procedure
var ColdRing *ring.Ring

// HotRingSize is default hot ring size
// Initialize with 10
var HotRingSize int

// ColdRingSize is default cold ring size
// Initialize with 10
var ColdRingSize int

// SoyLogExplorer rpc server client
var SoyLogExplorer client.XClient

// ExplorerServer is soy-log-explorer ip address
var ExplorerServer string

// ExplorerAddr is flag for ExplorerServer
var ExplorerAddr *string

func init() {
	HotRingSize = 10
	ColdRingSize = 10

	MapTable = make(map[string][]string)
	HotRing = ring.New(HotRingSize)
	ColdRing = ring.New(ColdRingSize)

	ExplorerServer = os.Getenv("EXPLORERSERVER")
	ExplorerAddr = flag.String("addr", ExplorerServer, "server address")

	SoyLogExplorer = CreateExplorerServer()
}
