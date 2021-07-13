package rpc

import (
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

// CopyLogMessage is deep copy with LogMessage structure
// Return copied LogMessage
func CopyLogMessage(arg *rpc.LogMessage) rpc.LogMessage {
	info := make([]rpc.LogInfo, len(arg.Info))
	buffer := make([]byte, len(arg.Buffer))
	copy(info, arg.Info)
	copy(buffer, arg.Buffer)

	ret := rpc.LogMessage{Info: info, Buffer: buffer}
	return ret
}
