package rpc

// CopyLogMessage is deep copy with LogMessage structure
// Return copied LogMessage
func CopyLogMessage(arg *LogMessage) LogMessage {
	info := make([]LogInfo, len(arg.info))
	buffer := make([]byte, len(arg.buffer))
	copy(info, arg.info)
	copy(buffer, arg.buffer)

	ret := LogMessage{info: info, buffer: buffer}
	return ret
}
