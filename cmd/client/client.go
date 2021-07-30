package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

var rpcserver = os.Getenv("RPCSERVER")
var addr = flag.String("addr1", rpcserver, "server address")

func main() {
	comp := &(compressor.GzipComp{})
	flag.Parse()

	var reply rpc.Reply
	var logmsg rpc.LogMessage

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")
	xclient := client.NewXClient("ColdPort", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	ixclient := client.NewXClient("Init", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	logmsg.Info = make([]rpc.LogInfo, 2)
	logmsg.Namespace = "application1:test"
	logmsg.Files.MapTable = append(logmsg.Files.MapTable, "File1")
	logmsg.Files.MapTable = append(logmsg.Files.MapTable, "File2")
	ixclient.Call(context.Background(), "Push", &logmsg, &reply)
	time.Sleep(10)
	logmsg.Files.Indexes = make([]uint8, 2)
	logmsg.Files.Indexes[0] = 1
	logmsg.Files.Indexes[1] = 0
	logmsg.Info[0].Timestamp = time.Now().UnixNano()
	logmsg.Info[0].Length = 9
	logmsg.Info[1].Timestamp = time.Now().UnixNano()
	logmsg.Info[1].Length = 4
	logmsg.Buffer, _ = comp.Compress([]byte("kijunkingkerr"))
	for {
		xclient.Call(context.Background(), "Push", &logmsg, &reply)
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
}
