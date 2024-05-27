package withtestmain_test

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.M) {
	fmt.Println("TestMain")
}

func TestDebug(t *testing.T) {
	fmt.Println("stopping on a breakpoint in this test works with delve")
}