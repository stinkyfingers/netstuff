package packets

import (
	"testing"
)

func TestRead(t *testing.T) {
	err := ReadFile("../cap.pcap")
	if err != nil {
		t.Error(err)
	}
}
