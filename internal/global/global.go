package global

import (
	"context"
	"flag"
	"os"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_collector/pkg/server"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// RedisServer is used for rpc data server with Hot/Cold port.
var RedisServer *server.Server

// Compressor is for rpc.
var Compressor compressor.Compressor

var ctx context.Context

// SoyLogExplorer rpc server client
var SoyLogExplorer client.XClient

// ExplorerServer is soy-log-explorer ip address
var ExplorerServer string

// ExplorerAddr is flag for ExplorerServer
var ExplorerAddr *string

// MapTable is used for log file indexes
var MapTable map[string][]string

func init() {
	ctx = context.Background()
	Compressor = &(compressor.GzipComp{})
	RedisServer = server.New(ctx)

	ExplorerServer = os.Getenv("EXPLORERSERVER")
	ExplorerAddr = flag.String("addr", ExplorerServer, "server address")

	SoyLogExplorer = CreateExplorerServer()

	MapTable = make(map[string][]string)
}
