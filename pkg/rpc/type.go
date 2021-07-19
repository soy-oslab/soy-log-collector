package rpc

// LogInfo is part of LogMessage.
// Consist of timestamp, filename, length.
// timestamp is the time log created.
// length is the size(bytes) per log message.
type LogInfo struct {
	Timestamp int64
	Length    uint64
}

// LogFile is part of LogMessage.
// Consist of maptable, indexes.
// maptable's key is filename and value is filename identifier.
// indexes contains filename identifier is defined MapTable.
type LogFile struct {
	MapTable []string
	Indexes  []uint8
}

// LogMessage is rpc parameter type with soy_log_generator.
// Consist of namespace, files, info, buffer.
// namespace is container's namespace.
// files contains the LogFile structure information/
// info is log information array.
// buffer is sequence of byte data.
// Compacted when used for ColdPort,
// Should not compacted when used for HotPort.
type LogMessage struct {
	Namespace string
	Files     LogFile
	Info      []LogInfo
	Buffer    []byte
}

// Reply is for Communication with rpc caller.
// Rate is current port utility.
type Reply struct {
	Rate uint8
}

// Move to pkg
