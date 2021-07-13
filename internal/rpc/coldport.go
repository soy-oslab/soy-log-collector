package rpc

import (
	"context"
	"errors"

	"github.com/soyoslab/soy_log_collector/internal/global"
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
	if global.ColdRing.Size() >= global.DefaultRingSize {
		reply.Rate = util.RangeMapping(global.ColdRing.Size())
		return errors.New("coldport is full")
	}
	log := CopyLogMessage(args)
	global.ColdRing.Push(&log)
	reply.Rate = util.RangeMapping(global.ColdRing.Size())
	return nil
}
