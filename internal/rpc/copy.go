package rpc

import (
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

// CopyLogMessage is deep copy with LogMessage structure
// Return copied LogMessage
func CopyLogMessage(arg *rpc.LogMessage) rpc.LogMessage {
	info := make([]rpc.LogInfo, len(arg.Info))
	buffer := make([]byte, len(arg.Buffer))
	var files rpc.LogFile

	files.MapTable = make([]string, len(arg.Files.MapTable))
	files.Indexes = make([]uint8, len(arg.Files.Indexes))
	copy(files.MapTable, arg.Files.MapTable)
	copy(files.Indexes, arg.Files.Indexes)
	buf := make([]byte, len(arg.Namespace))
	copy(buf, arg.Namespace)
	namespace := string(buf)
	copy(info, arg.Info)
	copy(buffer, arg.Buffer)
	ret := rpc.LogMessage{Namespace: namespace, Files: files, Info: info, Buffer: buffer}
	return ret
}
