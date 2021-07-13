package global

import (
	"testing"
)

func TestInit(t *testing.T) {
	if HotRing == nil {
		t.Error("HotRing is not created")
	}

	if ColdRing == nil {
		t.Error("ColdRing is not created")
	}

	if RedisServer == nil {
		t.Error("RedisServer is not created")
	}

	if Compressor == nil {
		t.Error("Compressor is not created")
	}

	if ctx == nil {
		t.Error("Context is not exist")
	}

	if DefaultRingSize == 0 {
		t.Error("Default Ring Size is not set")
	}
}
