package server

import (
	"flag"
	"os"

	"github.com/smallnest/rpcx/server"
	bg "github.com/soyoslab/soy_log_collector/internal/background"
	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
	daemon "github.com/soyoslab/soy_log_explorer/pkg/deamon"
)

var rpcserver = os.Getenv("RPCSERVER")
var addr = flag.String("addr", rpcserver, "server address")

// Server run rpcx server
func Server() {
	go daemon.Listen(global.HotRing, bg.HotPortHandler, 1)
	go daemon.Listen(global.ColdRing, bg.ColdPortHandler, 2)

	s := server.NewServer()
	err := s.RegisterName("HotPort", new(rpc.HotPort), "")
	if err != nil {
		return
	}
	err = s.RegisterName("ColdPort", new(rpc.ColdPort), "")
	if err != nil {
		return
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		return
	}
}
