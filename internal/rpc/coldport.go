package rpc

import (
	"context"
	"errors"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/util"
)

// ColdPort is rpc procedure type.
// Data should be compacted.
type ColdPort int

// Push insert data from rpc call into ring buffer.
// Return error when ColdPort is full.
// Communicate with caller via reply.
// Send current ColdPort utility with reply
func (t *ColdPort) Push(ctx context.Context, args *LogMessage, reply *Reply) error {
	if global.ColdRing.Size() >= global.DefaultRingSize {
		reply.Rate = util.RangeMapping(global.ColdRing.Size())
		return errors.New("ColdPort is full!")
	}
	log := CopyLogMessage(args)
	global.ColdRing.Push(&log)
	reply.Rate = util.RangeMapping(global.ColdRing.Size())
	return nil
}
