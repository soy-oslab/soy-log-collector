package rpc

// CopyLogMessage is deep copy with LogMessage structure
// Return copied LogMessage
func CopyLogMessage(arg *LogMessage) LogMessage {
	info := make([]LogInfo, len(arg.Info))
	buffer := make([]byte, len(arg.Buffer))
	copy(info, arg.Info)
	copy(buffer, arg.Buffer)

	ret := LogMessage{Info: info, Buffer: buffer}
	return ret
}
