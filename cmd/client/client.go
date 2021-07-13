package main

import (
	"context"
	"flag"
	"os"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_collector/pkg/rpc"
)

var rpcserver = os.Getenv("RPCSERVER")
var addr = flag.String("addr", rpcserver, "server address")

func main() {
	flag.Parse()

	var reply rpc.Reply
	var logmsg rpc.LogMessage

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("HotPort", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	logmsg.Info = make([]rpc.LogInfo, 2)
	logmsg.Info[0].Timestamp = time.Now().UnixNano()
	logmsg.Info[0].Filename = "test file"
	logmsg.Info[0].Length = 5
	logmsg.Info[1].Timestamp = time.Now().UnixNano()
	logmsg.Info[1].Filename = "bbbbbb"
	logmsg.Info[1].Length = 3
	logmsg.Buffer = []byte("hellobye")

	xclient.Call(context.Background(), "Push", &logmsg, &reply)
}
