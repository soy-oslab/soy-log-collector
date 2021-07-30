package main

import (
	"flag"
	"os"
	"runtime"

	"github.com/smallnest/rpcx/server"
	bg "github.com/soyoslab/soy_log_collector/internal/background"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
	daemon "github.com/soyoslab/soy_log_explorer/pkg/daemon"
)

// daemon polling time
var rpcserver = os.Getenv("RPCSERVER")
var addr = flag.String("rpcaddr", rpcserver, "server address")

// Server run rpcx server
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go daemon.Listen(rpc.HotRing, bg.HotPortHandler, 0)
	go daemon.Listen(rpc.ColdRing, bg.ColdPortHandler, 20)

	s := server.NewServer()
	err := s.RegisterName("HotPort", new(rpc.HotPort), "")
	if err != nil {
		return
	}
	err = s.RegisterName("ColdPort", new(rpc.ColdPort), "")
	if err != nil {
		return
	}
	err = s.RegisterName("Init", new(rpc.Init), "")
	if err != nil {
		return
	}

	err = s.Serve("tcp", *addr)
	if err != nil {
		return
	}
}
