package rpc

import (
	"context"
	"errors"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/util"
)

// HotPort is rpc procedure type.
// No compaction allowed for data.
type HotPort int

// Push insert data from rpc call into ring buffer.
// Return error when HotPort is full.
// Communicate with caller via reply.
// Send current HotPort utility with reply
func (t *HotPort) Push(ctx context.Context, args *LogMessage, reply *Reply) error {
	if global.HotRing.Size() >= global.DefaultRingSize {
		reply.Rate = util.RangeMapping(global.HotRing.Size())
		return errors.New("HotPort is full!")
	}
	log := CopyLogMessage(args)
	global.HotRing.Push(&log)
	reply.Rate = util.RangeMapping(global.HotRing.Size())
	return nil
}
