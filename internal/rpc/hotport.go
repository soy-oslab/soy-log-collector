package rpc

import (
	"context"
	"fmt"

	"github.com/soyoslab/soy_log_collector/internal/util"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
	"github.com/soyoslab/soy_log_explorer/pkg/signal"
)

// HotPort is rpc procedure type.
// No compaction allowed for data.
type HotPort int

// Push insert data from rpc call into ring buffer.
// Return error when HotPort is full.
// Communicate with caller via reply.
// Send current HotPort utility with reply
func (t *HotPort) Push(ctx context.Context, args *rpc.LogMessage, reply *rpc.Reply) error {
	err := checkAvailable(1)
	fmt.Println("Receive-Hot: ", args)
	if err != nil {
		reply.Rate = util.RangeMapping(HotRing.Size(), HotRingSize)
		return err
	}

	log := CopyLogMessage(args)
	HotRing.Push(&log)
	signal.Signal()
	reply.Rate = util.RangeMapping(HotRing.Size(), HotRingSize)
	return nil
}
