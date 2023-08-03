package main

import "testing"

func TestAdd(t *testing.T) {
	res, ctl := Add(12, 4), (12 + 4)
	if res != ctl {
		t.Errorf("\nIncorrect result, \n\tExpected: %v\n\tGot: %v\n", ctl, res)
	}
}

func TestSub(t *testing.T) {
	res, ctl := Subtract(12, 4), (12 - 4)
	if res != ctl {
		t.Errorf("\nIncorrect result, \n\tExpected: %v\n\tGot: %v\n", ctl, res)
	}
}

func TestDoMath(t *testing.T) {
	res, ctl := DoMath(12, 4, Add), (12 + 4)
	if res != ctl {
		t.Errorf("\nIncorrect result, \n\tExpected: %v\n\tGot: %v\n", ctl, res)
	}

	res, ctl = DoMath(12, 4, Subtract), (12 - 4)
	if res != ctl {
		t.Errorf("\nIncorrect result, \n\tExpected: %v\n\tGot: %v\n", ctl, res)
	}
}
