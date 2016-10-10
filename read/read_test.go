package read

import (
	"os"
	"testing"
	"time"

	"github.com/stinkyfingers/netstuff/device"
	"github.com/stinkyfingers/netstuff/write"
)

var name = "test.pcap"

func TestRead(t *testing.T) {
	resp, err := Read(name)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(resp))
}

func setup() error {
	var snapshotLen uint32 = 1024
	var err error
	go func() {
		packetSource, err := device.OpenDevice("en0")
		if err != nil {
			return
		}
		err = write.Write(name, snapshotLen, *packetSource)
		if err != nil {
			return
		}
	}()
	time.Sleep(time.Second * 3)
	return err
}

func teardown() error {
	return os.Remove(name)
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		return
	}
	c := m.Run()
	teardown()
	defer os.Exit(c)

}
