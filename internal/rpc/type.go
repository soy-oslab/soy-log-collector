package rpc

import (
	"time"
)

// LogInfo is part of LogMessage.
// Consist of timestamp, filename, length.
// timestamp is the time log created.
// filename is the name of log file.
// length is the length of log message.
type LogInfo struct {
	timestamp time.Time
	filename  string
	length    uint64
}

// LogMessage is rpc parameter type with soy_log_generator.
// Consist of info, buffer.
// info is log information array.
// buffer is sequence of byte data.
// Compacted when used for ColdPort,
// Should not compacted when used for HotPort.
type LogMessage struct {
	info   []LogInfo
	buffer []byte
}

// Reply is for Communication with rpc caller.
// Rate is current port utility.
type Reply struct {
	Rate uint8
}
