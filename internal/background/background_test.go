package background

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/internal/rpc"
)

func TestHotPortHandler(t *testing.T) {
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

	HotPortHandler(&logmsg)
}

func TestColdPortHandler(t *testing.T) {
	var err error
	var msg string
	var buffer []byte
	var logmsg rpc.LogMessage

	msgCount := 100

	logmsg.Info = make([]rpc.LogInfo, msgCount)
	for i := 0; i < msgCount; i++ {
		logmsg.Info[i].Timestamp = time.Now().UnixNano()
		logmsg.Info[i].Filename = "Cold Port Handler Test" + strconv.Itoa(i)
		msg = strconv.Itoa(rand.Int())
		logmsg.Info[i].Length = uint64(len(msg))
		buffer = append(buffer, []byte(msg)...)
	}

	logmsg.Buffer, err = global.Compressor.Compress(buffer)

	if err != nil {
		t.Error(err)
	}

	ColdPortHandler(&logmsg)
}
