package main

import "fmt"

func DoSomething(i, j int) (result int, err error) {

	defer func() {
		if err != nil {
			err = fmt.Errorf("in DoSomething: %w", err)
		}
	}()
	err = doThing1()
	if err != nil {
		return 0, err
	}
	err = doThing2()
	if err != nil {
		return 0, err
	}
	err = doThing3()
	if err != nil {
		return 0, err
	}

	return i + j, nil
}

func doThing1() error {
	return nil
}
func doThing2() error {
	return nil
}
func doThing3() error {
	return nil
}
