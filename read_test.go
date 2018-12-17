package go_1w

import (
	//	"github.com/goneri/go_1w"
	"testing"
)

func TestRead(t *testing.T) {
	value := Read("sample")
	expectation := float32(-19.625)
	if value != expectation {
		t.Errorf("Incorrect value: %f, want: %f.", value, expectation)
	}
}
