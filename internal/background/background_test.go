package background

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
	irpc "github.com/soyoslab/soy_log_collector/internal/rpc"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

func setMapTable() {
	var table []string
	table = append(table, "File1")
	table = append(table, "File2")
	irpc.MapTable["TestModule:test"] = table
}

func makeMsg(hotcold bool) rpc.LogMessage {
	var err error
	var msg string
	var logmsg rpc.LogMessage
	msgCount := 100

	logmsg.Info = make([]rpc.LogInfo, msgCount)
	logmsg.Namespace = "TestModule:test"
	logmsg.Files.Indexes = make([]uint8, 100)
	for i := 0; i < msgCount; i++ {
		logmsg.Info[i].Timestamp = time.Now().UnixNano()
		msg = strconv.Itoa(rand.Int())
		logmsg.Info[i].Length = uint64(len(msg))
		logmsg.Files.Indexes[i] = 0
		logmsg.Buffer = append(logmsg.Buffer, []byte(msg)...)
	}

	if !hotcold {
		logmsg.Buffer, err = global.Compressor.Compress(logmsg.Buffer)
		if err != nil {
			panic(err)
		}
	}

	return logmsg
}

func TestHotPortHandler(t *testing.T) {
	setMapTable()
	logmsg := makeMsg(true)

	HotPortHandler(&logmsg)
}

func TestColdPortHandler(t *testing.T) {
	setMapTable()
	logmsg := makeMsg(false)

	ColdPortHandler(&logmsg)
}
