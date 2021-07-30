package global

import (
	"testing"
)

func TestInit(t *testing.T) {
	if RedisServer == nil {
		t.Error("RedisServer is not created")
	}

	if Compressor == nil {
		t.Error("Compressor is not created")
	}

	if ctx == nil {
		t.Error("Context is not exist")
	}
}
