package background

import (
	"context"

	decoder "github.com/mitchellh/mapstructure"
	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/util"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

type esdocs struct {
	Index string
	Docs  string
}

// HotPortHandler is processing unit with HotRing.
// Running for goroutine with daemon.
func HotPortHandler(args ...interface{}) {
	var buf *rpc.LogMessage
	err := decoder.Decode(args[0], &buf)

	if err != nil {
		panic(err)
	}

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
		date, sec, nano, err = util.TimeSlice(timestamp)
		if err != nil {
			panic(err)
		}
		key = filename + ":" + date + ":" + sec + ":" + nano
		err = global.RedisServer.Push(key, log)
		if err != nil {
			panic(err)
		}
		idx += length
		sendMessage(key, log, true)
	}

}

func sendMessage(idx string, data string, hotcold bool) {
	var docs esdocs
	var reply string

	docs.Index = idx
	docs.Docs = data
	if hotcold {
		global.SoyLogExplorer.Call(context.Background(), "HotPush", &docs, &reply)
	} else {
		global.SoyLogExplorer.Call(context.Background(), "ColdPush", &docs, &reply)
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
		panic(err)
	}

	for _, loginfo := range buf.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = loginfo.Filename
		date, sec, nano, err = util.TimeSlice(timestamp)
		if err != nil {
			panic(err)
		}
		key = filename + ":" + date + ":" + sec + ":" + nano
		log = string(buffer[idx : idx+length])
		err = global.RedisServer.Push(key, log)
		if err != nil {
			panic(err)
		}
		idx += length
	}
}
