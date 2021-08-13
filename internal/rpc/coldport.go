package rpc

import (
	"context"
	"fmt"

	"github.com/soyoslab/soy_log_collector/internal/util"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
	"github.com/soyoslab/soy_log_explorer/pkg/signal"
)

// ColdPort is rpc procedure type.
// Data should be compacted.
type ColdPort int

// Push insert data from rpc call into ring buffer.
// Return error when ColdPort is full.
// Communicate with caller via reply.
// Send current ColdPort utility with reply
func (t *ColdPort) Push(ctx context.Context, args *rpc.LogMessage, reply *rpc.Reply) error {
	err := checkAvailable(0)
	fmt.Println("Receive-Cold: ", args)
	if err != nil {
		reply.Rate = util.RangeMapping(ColdRing.Size(), ColdRingSize)
		return err
	}

	log := CopyLogMessage(args)
	ColdRing.Push(&log)
	signal.Signal()
	reply.Rate = util.RangeMapping(ColdRing.Size(), ColdRingSize)
	return nil
}
