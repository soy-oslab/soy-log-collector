package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
)

func TestHotPush(t *testing.T) {
	ctx := context.Background()

	var hotport HotPort
	var logmsg LogMessage
	var reply Reply

	logmsg.Info = make([]LogInfo, 1)
	logmsg.Info[0].Timestamp = time.Now().UnixNano()
	logmsg.Info[0].Filename = "Hot Push Test"
	logmsg.Info[0].Length = 5
	logmsg.Buffer = []byte("hello")

	err := hotport.Push(ctx, &logmsg, &reply)

	if err != nil {
		t.Error(err)
	}
}

func TestColdPush(t *testing.T) {
	ctx := context.Background()

	var coldport ColdPort
	var logmsg LogMessage
	var reply Reply
	var err error

	logmsg.Info = make([]LogInfo, 1)
	logmsg.Info[0].Timestamp = time.Now().UnixNano()
	logmsg.Info[0].Filename = "Hot Push Test"
	logmsg.Info[0].Length = 5
	logmsg.Buffer, err = global.Compressor.Compress([]byte("hello"))

	if err != nil {
		t.Error(err)
	}

	err = coldport.Push(ctx, &logmsg, &reply)

	if err != nil {
		t.Error(err)
	}
}
