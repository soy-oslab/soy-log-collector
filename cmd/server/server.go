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

var addr = flag.String("addr", os.Getenv("RPCSERVER"), "server address")

// Server run rpcx server
func Server() {
	s := server.NewServer()
	err := s.Register(new(rpc.HotPort), "")
	if err != nil {
		return
	}
	err = s.Register(new(rpc.ColdPort), "")
	if err != nil {
		return
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		return
	}

	go daemon.Listen(global.HotRing, bg.HotPortHandler, 1)
	go daemon.Listen(global.ColdRing, bg.ColdPortHandler, 2)
}
