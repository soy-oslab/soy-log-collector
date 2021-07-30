package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"flag"
	"fmt"

	"github.com/smallnest/rpcx/client"
	"github.com/soyoslab/soy_log_generator/pkg/compressor"
)

var addr = flag.String("addr", "164.125.68.227:8972", "server address")

type esdocs struct {
	Index string
	Docs  string
}

func main() {
	flag.Parse()

	var reply string

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*addr, "")

	xclient := client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	docs := esdocs{Index: "hothothot", Docs: `{"name":"kai"}`}
	var buf bytes.Buffer

	enc := gob.NewEncoder(&buf)
	err := enc.Encode(docs)
	if err != nil {
		fmt.Printf("error1: %v\n", err)
		return
	}

	c := &compressor.GzipComp{}
	data, err := c.Compress(buf.Bytes())
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	err = xclient.Call(context.Background(), "HotPush", &data, &reply)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("reply: %s\n", reply)
	}
}
