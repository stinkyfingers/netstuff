package device

import (
	"testing"
)

func TestDevices(t *testing.T) {
	devices, err := Devices()
	if err != nil {
		t.Error(err)
	}
	if len(devices) < 1 {
		t.Error("Expected devices.")
	}
}

func TestOpenDevice(t *testing.T) {
	name := "en0"
	_, err := OpenDevice(name)
	if err != nil {
		t.Error(err)
	}
}
