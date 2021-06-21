package main

import (
    "github.com/soyoslab/soy_log_collector/soy/lib"
	"github.com/smallnest/rpcx/server"
	"flag"
)

var addr = flag.String("addr", "localhost:8972", "server address")

func main() {
	s := server.NewServer()
	s.Register(new(HotPort), "")
	s.Serve("tcp", *addr)
}
