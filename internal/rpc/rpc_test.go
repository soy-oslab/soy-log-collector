package rpc

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/soyoslab/soy_log_collector/internal/global"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

func setMapTable() []string {
	var table []string

	table = append(table, "File1")
	table = append(table, "File2")
	MapTable["TestModule:test"] = table

	return table
}

func makeMsg(hotcold bool) rpc.LogMessage {
	var err error
	var msg string
	var logmsg rpc.LogMessage
	msgCount := 100

	logmsg.Info = make([]rpc.LogInfo, msgCount)
	logmsg.Namespace = "TestModule:test"
	logmsg.Files.Indexes = make([]uint8, msgCount)

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

func TestInitPush(t *testing.T) {
	var logmsg rpc.LogMessage
	var reply rpc.Reply
	var initport Init

	ctx := context.Background()

	logmsg.Namespace = "app1"
	logmsg.Files.MapTable = setMapTable()

	initport.Push(ctx, &logmsg, &reply)
}

func TestInitPushFail(t *testing.T) {
	var logmsg rpc.LogMessage
	var reply rpc.Reply
	var initport Init
	var table []string
	ctx := context.Background()

	logmsg.Namespace = "app1"
	logmsg.Files.MapTable = table
	initport.Push(ctx, &logmsg, &reply)
}

func TestCheckAvailable(t *testing.T) {
	ctx := context.Background()

	var hotport HotPort
	var reply rpc.Reply

	logmsg := makeMsg(true)
	InitFlag = 0
	hotport.Push(ctx, &logmsg, &reply)
}

func TestHotPush(t *testing.T) {
	ctx := context.Background()

	var hotport HotPort
	var reply rpc.Reply

	logmsg := makeMsg(true)
	setMapTable()
	InitFlag = 1
	for i := 0; i < 10; i++ {
		err := hotport.Push(ctx, &logmsg, &reply)
		if err != nil {
			t.Error(err)
		}
	}

	hotport.Push(ctx, &logmsg, &reply)
}

func TestColdPush(t *testing.T) {
	ctx := context.Background()

	var coldport ColdPort
	var reply rpc.Reply
	setMapTable()
	InitFlag = 1

	logmsg := makeMsg(false)

	for i := 0; i < 10; i++ {
		err := coldport.Push(ctx, &logmsg, &reply)
		if err != nil {
			t.Error(err)
		}
	}

	coldport.Push(ctx, &logmsg, &reply)
}
