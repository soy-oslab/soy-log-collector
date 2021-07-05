package server

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

// Server structure with multiple db name
type Server struct {
	db  redis.Client
	ctx context.Context
}

// New return Server structure with initializing DBTable
func New(ctx context.Context) Server {
	s := Server{}
	s.db = s.newClient(0)
	s.ctx = context.Background()

	return s
}

func (t *Server) newClient(dbnum int) redis.Client {
	addr := os.Getenv("DBADDR")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       dbnum,
	})

	return *rdb
}

// Push insert datas into db
// db select db want to insert {key, data}
// key matching to data
// data matching to key
func (t *Server) Push(key string, data string) {
	ctx := t.ctx
	err := t.db.Set(ctx, key, data, 0).Err()
	if err != nil {
		panic(err)
	}
}
