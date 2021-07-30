package rpc

import (
	"context"
	"errors"

	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

// Init is rpc procedure type.
// Set MapTable.
type Init int

// Push insert data from rpc call into ring buffer.
// Return error when ColdPort is full.
// Communicate with caller via reply.
// Send current ColdPort utility with reply
func (t *Init) Push(ctx context.Context, args *rpc.LogMessage, reply *rpc.Reply) error {
	if len(args.Files.MapTable) < 1 {
		reply.Rate = 0
		return errors.New("no maptables")
	}
	MapTable[args.Namespace] = args.Files.MapTable
	InitFlag = 1
	return nil
}
