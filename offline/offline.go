package offline

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	packetHandle "github.com/stinkyfingers/netstuff/packet"
)

func Open(pcapFile string) error {
	handle, err := pcap.OpenOffline(pcapFile)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		err = packetHandle.HandlePacket(packet)
		if err != nil {
			if err.Error() == "Content not IPv4" {
				continue
			}
			return err
		}
	}
	return err
}
