package server

import (
	"context"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	server := New(ctx)

	fmt.Println(ctx)
	fmt.Println(server.ctx)
}

func TestPush(t *testing.T) {
	ctx := context.Background()
	server := New(ctx)

	err := server.Push("key", "value")

	if err != nil {
		t.Error(err)
	}
}
