package background

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	decoder "github.com/mitchellh/mapstructure"
	"github.com/soyoslab/soy_log_collector/internal/global"
	irpc "github.com/soyoslab/soy_log_collector/internal/rpc"
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
func SendMessage(idx string, data string, isHot bool) error {
	var docs esdocs.ESdocs
	var reply string
	var err error
	var compressed []byte
	var connectionCount int

	connectionCount = 0
	docs.Index = idx
	docs.Docs = data
	err = errors.New("full")

	if !isHot {
		compressed, _ = docsCompress(docs)
	}

SEND:
	if err != nil {
		if strings.Contains(err.Error(), "full") {
			for err != nil {
				if isHot {
					err = irpc.SoyLogExplorer.Call(context.Background(), "HotPush", &docs, &reply)
				} else {
					err = irpc.SoyLogExplorer.Call(context.Background(), "ColdPush", &compressed, &reply)
				}
			}
		} else {
			if connectionCount > 10 {
				return errors.New("connection fail")
			}
			irpc.SoyLogExplorer = irpc.CreateExplorerServer()
			connectionCount++
		}
		goto SEND
	}
	fmt.Println("Send-", isHot, ": ", docs)
	return nil
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

func makeJSON(timestamp string, filename string, log string, host string) map[string]string {
	buf := make(map[string]string)

	buf["@timestamp"] = timestamp
	buf["@filename"] = filename
	buf["log"] = log
	buf["@host"] = host

	return buf
}

func handler(arg *rpc.LogMessage, isHot bool) {
	var (
		idx                            int
		err                            error
		length                         int
		log                            string
		timestamp                      int64
		logarr                         []map[string]string
		filename, key, date, sec, nano string
		namespace                      string
		host                           string
	)

	logarr = make([]map[string]string, 0)
	ns := strings.Split(arg.Namespace, ":")
	namespace = ns[0]
	host = ns[1]
	idx = 0

	for i, loginfo := range arg.Info {
		length = int(loginfo.Length)
		timestamp = loginfo.Timestamp
		filename = irpc.MapTable[arg.Namespace][arg.Files.Indexes[i]]
		log = string(arg.Buffer[idx : idx+length])
		ts := util.TimeSlice(timestamp)
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
			logarr = logarr[:0]
			logarr = append(logarr, makeJSON(ts, filename, log, host))
			jsonfy, _ := json.Marshal(logarr)
			SendMessage(namespace, string(jsonfy), isHot)
		} else {
			err := Filter(log)
			if err != nil {
				continue
			}
			logarr = append(logarr, makeJSON(ts, filename, log, host))
		}
	}
	if !isHot && len(logarr) > 0 {
		jsonfy, _ := json.Marshal(logarr)
		SendMessage(namespace, string(jsonfy), isHot)
	}
}
