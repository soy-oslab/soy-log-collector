package background

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

func makeMsg(hotcold bool) rpc.LogMessage {
	var err error
	var msg string
	var logmsg rpc.LogMessage
	msgCount := 100

	logmsg.Info = make([]rpc.LogInfo, msgCount)
	for i := 0; i < msgCount; i++ {
		logmsg.Info[i].Timestamp = time.Now().UnixNano()
		logmsg.Info[i].Filename = "Hot Port Handler Test" + strconv.Itoa(i)
		msg = strconv.Itoa(rand.Int())
		logmsg.Info[i].Length = uint64(len(msg))
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
	logmsg := makeMsg(true)

	HotPortHandler(&logmsg)
}

func TestColdPortHandler(t *testing.T) {
	logmsg := makeMsg(false)

	ColdPortHandler(&logmsg)
}
