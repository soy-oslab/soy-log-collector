package background

import (
	decoder "github.com/mitchellh/mapstructure"
	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
	"github.com/soyoslab/soy_log_collector/internal/util"
)

// HotPortHandler is processing unit with HotRing.
// Running for goroutine with daemon.
func HotPortHandler(args ...interface{}) {
	var buf *rpc.LogMessage
	decoder.Decode(args[0], &buf)

	var timestamp int64
	var idx int
	var length int
	var log string
	var filename, key, date, sec, nano string

	idx = 0
	for _, loginfo := range buf.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = loginfo.Filename
		log = string(buf.Buffer[idx : idx+length])
		date, sec, nano = util.TimeSlice(timestamp)
		key = filename + ":" + date + ":" + sec + ":" + nano
		global.RedisServer.Push(key, log)
		idx += length
	}
}

// ColdPortHandler is processing unit with ColdRing.
// Running for goroutine with daemon.
func ColdPortHandler(args ...interface{}) {
	var buf *rpc.LogMessage
	decoder.Decode(args[0], &buf)

	var timestamp int64
	var idx int
	var length int
	var log string
	var filename, key, date, sec, nano string

	idx = 0
	buffer, err := global.Compressor.Decompress(buf.Buffer)

	if err != nil {
		panic("compress failed")
	}

	for _, loginfo := range buf.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = loginfo.Filename
		date, sec, nano = util.TimeSlice(timestamp)
		key = filename + ":" + date + ":" + sec + ":" + nano
		log = string(buffer[idx : idx+length])
		global.RedisServer.Push(key, log)
		idx += length
	}
}
