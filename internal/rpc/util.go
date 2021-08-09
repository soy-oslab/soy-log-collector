package rpc

import "errors"

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
	if hotcold == 1 {
		if HotRing.Size() >= HotRingSize {
			return errors.New("hotport is full")
		}
	} else {
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
