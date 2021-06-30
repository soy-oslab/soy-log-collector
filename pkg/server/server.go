package server

import (
	"context"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

// Server structure with multiple db name
type Server struct {
	dblist map[string]redis.Client
	ctx    context.Context
}

// New return Server structure with initializing DBTable
func New(ctx context.Context) Server {
	dblist := make(map[string]redis.Client)
	s := Server{dblist: dblist, ctx: ctx}
	s.dblist["root"] = s._NewClient(0)

	s._InitDBTable()

	return s
}

func (t *Server) _NewClient(dbnum int) redis.Client {
	addr := os.Getenv("DBADDR")
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       dbnum,
	})

	return *rdb
}

func (t *Server) _InitDBTable() {
	dbtable := t.dblist["root"]
	ctx := t.ctx

	keys := dbtable.Keys(t.ctx, "").Val()
	for _, value := range keys {
		dbnum := dbtable.Get(ctx, value).Val()
		s, _ := strconv.ParseInt(dbnum, 10, 64)
		t.dblist[value] = t._NewClient(int(s))
	}
}

// Push insert datas into db
// db select db want to insert {key, data}
// key matching to data
// data matching to key
func (t *Server) Push(db string, key string, data string) {
	ctx := t.ctx

	if _, ok := t.dblist[db]; !ok {
		t.Push("root", db, strconv.Itoa(len(t.dblist)))
		t.dblist[key] = t._NewClient(len(t.dblist))
	}

	err := t.dblist[db].Set(ctx, key, data, 0).Err()
	if err != nil {
		panic(err)
	}
}
