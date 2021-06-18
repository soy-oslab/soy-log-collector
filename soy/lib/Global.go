package Global

import (
	"container/ring"
	"github.com/go-redis/redis/v8"
)

var HotRing ring
var ColdRing ring
var RingSize int
var rdb *redis.Client

func init() {
	RingSize = 10
	HotRing = ring.New(RingSize)
	ColdRing = ring.New(RingSize)
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
