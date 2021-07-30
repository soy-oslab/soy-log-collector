package global

import (
	"context"

	"github.com/soyoslab/soy_log_collector/pkg/server"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// RedisServer is used for rpc data server with Hot/Cold port.
var RedisServer *server.Server

// Compressor is for rpc.
var Compressor compressor.Compressor

var ctx context.Context

func init() {
	ctx = context.Background()
	Compressor = &(compressor.GzipComp{})
	RedisServer = server.New(ctx)
}
