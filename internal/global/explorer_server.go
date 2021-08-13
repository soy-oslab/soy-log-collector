package global

import (
	"github.com/smallnest/rpcx/client"
)

// CreateExplorerServer return soy_log_explorer rpc server client
func CreateExplorerServer() client.XClient {
	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*ExplorerAddr, "")
	return client.NewXClient("Rpush", client.Failtry, client.RandomSelect, d, client.DefaultOption)
}
