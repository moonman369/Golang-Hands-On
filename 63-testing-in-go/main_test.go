// filename should always be <file_to_be_tested>_test.go
package main

import "testing"

// Test funtion naming convention: Test<FuncName>(t *testing.T) {}
func TestAdd(t *testing.T) {
	sum, ctl := Add(5, 5), 5+5
	if sum != ctl {
		t.Errorf("\nIncorrect sum, \n\tExpected: %v\n\tGot: %v\n", ctl, sum)
	}
}
