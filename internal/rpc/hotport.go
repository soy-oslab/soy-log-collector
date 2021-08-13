package rpc

import (
	"context"
	"fmt"

	"github.com/soyoslab/soy_log_collector/internal/background"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
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
		return err
	}

	log := CopyLogMessage(args)
	go background.Handler(&log, true)
	return nil
}
