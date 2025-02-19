package my_testing

import (
	"fmt"
	"testing"
)

func Check_For_Panic(t *testing.T, should_panic bool) {
	r := recover()
	did_panic := r != nil
	if did_panic && !should_panic {
		fmt.Print("did incorrectly panic")
		t.Error()
	}
	if !did_panic && should_panic {
		fmt.Print("did not panic but should have")
		t.Error()
	}
}
