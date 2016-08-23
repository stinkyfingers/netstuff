package packets

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

func Offline(name string) error {
	// Open file instead of device
	handle, err = pcap.OpenOffline(name)
	if err != nil {
		return err
	}
	defer handle.Close()

	// Loop through packets in file
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		log.Println(packet)
	}
	return err
}
