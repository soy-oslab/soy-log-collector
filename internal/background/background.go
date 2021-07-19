package background

import (
	"bytes"
	"context"
	"encoding/gob"

	decoder "github.com/mitchellh/mapstructure"
	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/util"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
	"github.com/soyoslab/soy_log_explorer/pkg/esdocs"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

// HotPortHandler is processing unit with HotRing.
// Running for goroutine with daemon.
func HotPortHandler(args ...interface{}) {
	var buf *rpc.LogMessage
	err := decoder.Decode(args[0], &buf)

	if err != nil {
		panic(err)
	}

	handler(buf, true)
}

func docsCompress(docs esdocs.ESdocs) ([]byte, error) {
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(docs)
	if err != nil {
		return nil, err
	}

	c := &compressor.GzipComp{}
	return c.Compress(buf.Bytes())
}

func sendMessage(idx string, data string, hotcold bool) {
	var docs esdocs.ESdocs
	var reply string

	docs.Index = idx
	docs.Docs = data
	if hotcold {
		global.SoyLogExplorer.Call(context.Background(), "HotPush", &docs, &reply)
	} else {
		data, _ := docsCompress(docs)
		global.SoyLogExplorer.Call(context.Background(), "ColdPush", &data, &reply)
	}
}

// ColdPortHandler is processing unit with ColdRing.
// Running for goroutine with daemon.
func ColdPortHandler(args ...interface{}) {
	var buf *rpc.LogMessage
	decoder.Decode(args[0], &buf)

	buffer, err := global.Compressor.Decompress(buf.Buffer)

	if err != nil {
		panic(err)
	}

	buf.Buffer = buffer

	handler(buf, false)
}

func mergeString(value []string) string {
	merged := ""
	for _, v := range value {
		merged += v
	}

	return merged
}

func handler(arg *rpc.LogMessage, hotcold bool) {
	var timestamp int64
	var idx int
	var length int
	var log string
	var filename, key, date, sec, nano string
	var err error

	idx = 0

	var buf map[string][]string
	var coldlog map[string][]string

	for _, loginfo := range arg.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = loginfo.Filename
		log = string(arg.Buffer[idx : idx+length])
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
		if hotcold {
			sendMessage(key, log, hotcold)
		} else {
			buf = Filter(log)
			coldlog = MergeMap(buf, coldlog)
		}
	}

	if !hotcold {
		for key, value := range coldlog {
			merged := mergeString(value)
			sendMessage(key, merged, hotcold)
		}
	}
}
