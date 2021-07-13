package ring

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	ring := New()
	fmt.Println(ring.capacity)

	ring = New(20)
	fmt.Println(ring.capacity)
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
		t.Errorf("Can't Push!")
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
			fmt.Println(v)
		}
	}

	_, err = ring.Pop()
	if err == nil {
		t.Errorf("Must be empty!")
	}
}

func TestSize(t *testing.T) {
	var err error

	ring := New()

	for i := 0; i < 10; i++ {
		err = ring.Push(i)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(ring.Size())
	}

}
