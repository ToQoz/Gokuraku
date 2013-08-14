package test_helpers

import (
	"fmt"
	"testing"
)

func AssertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Error(fmt.Sprintf("expected <%s>, but got <%s>", expected, actual))
	}
}
