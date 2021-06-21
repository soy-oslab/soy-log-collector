package Service

import (
	"context"
	"fmt"
	"flag"
	lib "github.com/soyoslab/soy_log_collector/soy/lib"
	"github.com/smallnest/rpcx/server"
)

func rangeMapping(arg int)  int {
	return math.Round(arg * 255/ lib.Global.RingSize)
}

type Reply struct {
	Rate    uint8
}

type ColdPort int

func (t *ColdPort) Push(ctx context.Context, args string, reply *Reply) error {
	reply.Rate = rangeMapping(lib.Global.ColdRing.Len())
	if reply.Rate > 255 {
		return -1
	}
	return nil
}

type HotPort int

func (t *HotPort) Push(ctx context.Context, args string, reply *Reply) error {
	reply.Rate = rangeMapping(lib.Global.HotRing.Len())
	if reply.Rate > 255 {
		return -1
	}
	return nil
}
