package withouttestmain

import (
	"fmt"
	"testing"
)

func TestDebug(t *testing.T) {
	fmt.Println("stopping on a breakpoint in this test works with delve")
}