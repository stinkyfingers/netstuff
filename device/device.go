package device

import (
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

const (
	snaplen int32         = 1024
	promisc bool          = true
	timeout time.Duration = -1
)

func Devices() ([]pcap.Interface, error) {
	return pcap.FindAllDevs()
}

func OpenDevice(name string) (*gopacket.PacketSource, error) {
	handle, err := pcap.OpenLive(name, snaplen, promisc, timeout)
	if err != nil {
		return nil, err
	}
	return gopacket.NewPacketSource(handle, handle.LinkType()), nil

}
