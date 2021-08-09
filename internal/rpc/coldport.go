package rpc

import (
	"context"

	"github.com/soyoslab/soy_log_collector/internal/util"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

// ColdPort is rpc procedure type.
// Data should be compacted.
type ColdPort int

// Push insert data from rpc call into ring buffer.
// Return error when ColdPort is full.
// Communicate with caller via reply.
// Send current ColdPort utility with reply
func (t *ColdPort) Push(ctx context.Context, args *rpc.LogMessage, reply *rpc.Reply) error {
	err := checkInit()
	if err != nil {
		reply.Rate = 0
		return err
	}

	err = checkRingSize(0)
	if err != nil {
		reply.Rate = util.RangeMapping(ColdRing.Size(), ColdRingSize)
		return err
	}

	log := CopyLogMessage(args)
	ColdRing.Push(&log)
	reply.Rate = util.RangeMapping(ColdRing.Size(), ColdRingSize)
	return nil
}
