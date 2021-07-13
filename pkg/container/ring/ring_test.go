package ring

import (
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	ring := New()
	log.Println(ring.capacity)

	ring = New(20)
	log.Println(ring.capacity)
}

func TestPush(t *testing.T) {
	var err error

	ring := New()

	for i := 0; i < 10; i++ {
		err = ring.Push(i)
		if err != nil {
			t.Error(err)
		}
	}

	err = ring.Push(1)

	if err == nil {
		t.Errorf("can't push!")
	}
}

func TestPop(t *testing.T) {
	var err error

	ring := New()

	for i := 0; i < 5; i++ {
		err = ring.Push(i)
		if err != nil {
			t.Error(err)
		}
	}

	for i := 0; i < 5; i++ {
		v, err := ring.Pop()
		if err != nil {
			t.Error(err)
		} else {
			log.Println(v)
		}
	}

	_, err = ring.Pop()
	if err == nil {
		t.Errorf("must be empty!")
	}
}

func TestSize(t *testing.T) {
	ring := New()

	ring.Push(1)
	log.Println(ring.Size())
	ring.Push(2)
	log.Println(ring.Size())
}
