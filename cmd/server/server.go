package server

import (
	"flag"

	"github.com/smallnest/rpcx/server"
)

var addr = flag.String("addr", "localhost:8972", "server address")

// Server run rpcx server
func Server() {
	s := server.NewServer()
	err := s.Register(new(int), "")
	if err != nil {
		return
	}
	err = s.Serve("tcp", *addr)
	if err != nil {
		return
	}
}
