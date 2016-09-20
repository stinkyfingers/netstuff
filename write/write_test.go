package write

import (
	"os"
	"testing"
	"time"

	"github.com/stinkyfingers/netstuff/device"
)

func TestWrite(t *testing.T) {
	name := "test.pcap"
	var snapshotLen uint32 = 1024
	go func() {
		packetSource, err := device.OpenDevice("en0")
		if err != nil {
			t.Error(err)
		}
		err = Write(name, snapshotLen, *packetSource)
		if err != nil {
			t.Error(err)
		}
	}()
	time.Sleep(time.Second * 2)

	os.Remove(name)
}
