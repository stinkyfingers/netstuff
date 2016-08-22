package packets

import (
	"testing"
)

func TestDevices(t *testing.T) {
	err := Devices()
	if err != nil {
		t.Error(err)
	}
}
