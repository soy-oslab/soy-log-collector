package background

import (
	"bytes"
	"context"
	"encoding/gob"
	"strings"

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

// SendMessage send esdocs type data to soy_log_explorer
func SendMessage(idx string, data string, isHot bool) {
	var docs esdocs.ESdocs
	var reply string

	docs.Index = idx
	docs.Docs = data
	if isHot {
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
	merged := "{"
	for _, v := range value {
		slices := strings.Split(v, ":")
		log := slices[len(slices)-1]
		timestamp := ""
		for _, t := range slices {
			timestamp += t
		}
		merged += "\""
		merged += timestamp + "\":"
		merged += log + "\","
	}

	merged = merged[:len(merged)-2] + "}"

	return merged
}

func handler(arg *rpc.LogMessage, isHot bool) {
	var idx int
	var err error
	var length int
	var log string
	var timestamp int64
	var coldlog []string
	var filename, key, date, sec, nano, buf string

	idx = 0

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
		if isHot {
			SendMessage("hot", log, isHot)
		} else {
			buf = Filter(key, log)
			coldlog = append(coldlog, buf)
		}
	}

	if !isHot {
		merged := mergeString(coldlog)
		SendMessage("cold", merged, isHot)
	}
}
