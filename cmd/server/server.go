package server

import (
	"flag"
	"os"

	"github.com/smallnest/rpcx/server"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
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
}
