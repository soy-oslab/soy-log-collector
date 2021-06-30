package background

import (
	"time"

	decoder "github.com/mitchellh/mapstructure"
	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
)

// HotPortHandler is processing unit with HotRing.
// Running for goroutine with daemon.
func HotPortHandler(args ...interface{}) {
	var buf rpc.LogMessage
	decoder.Decode(args[0], &buf)

	var timestamp time.Time
	var idx int
	var length int
	var log string
	var filename string

	for _, loginfo := range buf.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = loginfo.Filename
		log = string(buf.Buffer[idx : idx+length])
		global.RedisServer.Push(filename, timestamp.String(), log)
	}
}
