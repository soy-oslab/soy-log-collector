package rpc

import (
	"errors"
)

func checkAvailable(hotcold int) error {
	err := checkInit()
	if err != nil {
		return err
	}

	err = checkRingSize(hotcold)
	if err != nil {
		return err
	}

	return nil
}

func checkRingSize(hotcold int) error {
	if hotcold == 0 {
		if ColdRing.Size() >= ColdRingSize {
			return errors.New("coldport is full")
		}
	}

	return nil
}

func checkInit() error {
	if InitFlag == 0 {
		return errors.New("init must be called first")
	}

	return nil
}
