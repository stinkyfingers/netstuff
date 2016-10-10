package layer

import (
	"testing"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	file = "../../../../../../Desktop/tpiscool/tpic2.pcap"
)

func TestLayer(t *testing.T) {
	// Open file instead of device
	handle, err := pcap.OpenOffline(file)
	if err != nil {
		t.Log(err)
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		err = Layer(packet)
		if err != nil {
			t.Error(err)
		}
	}
}
