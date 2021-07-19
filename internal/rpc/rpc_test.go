package rpc

import (
	"context"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

func makeMsg(hotcold bool) rpc.LogMessage {
	var logmsg rpc.LogMessage
	var err error

	buffer := []byte("hello")

	logmsg.Info = make([]rpc.LogInfo, 1)
	logmsg.Info[0].Timestamp = time.Now().UnixNano()
	logmsg.Info[0].Filename = "Hot Push Test"
	logmsg.Info[0].Length = 5

	if !hotcold {
		logmsg.Buffer, err = global.Compressor.Compress(buffer)
		if err != nil {
			panic(err)
		}
	} else {
		logmsg.Buffer = buffer
	}

	return logmsg
}

func TestHotPush(t *testing.T) {
	ctx := context.Background()

	var hotport HotPort
	var reply rpc.Reply

	logmsg := makeMsg(true)

	for i := 0; i < 10; i++ {
		err := hotport.Push(ctx, &logmsg, &reply)
		if err != nil {
			t.Error(err)
		}
	}

	err := hotport.Push(ctx, &logmsg, &reply)
	if err == nil {
		t.Errorf("hotport must be full")
	}

}

func TestColdPush(t *testing.T) {
	ctx := context.Background()

	var coldport ColdPort
	var reply rpc.Reply

	logmsg := makeMsg(false)

	for i := 0; i < 10; i++ {
		err := coldport.Push(ctx, &logmsg, &reply)
		if err != nil {
			t.Error(err)
		}
	}

	err := coldport.Push(ctx, &logmsg, &reply)
	if err == nil {
		t.Errorf("coldport must be full")
	}
}
