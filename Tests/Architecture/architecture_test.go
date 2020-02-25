package architecture

import (
	"runtime"
	"testing"
)

func TestDependent(t *testing.T) {
	t.Parallel()                   // can be executed in parallel
	if runtime.GOARCH == "amd64" { // test build type
		t.Skip("Doesn't work in architecture amd64")
	}

	// ...
	t.Fail()
}
