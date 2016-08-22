package packets

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Offline(path string) error {
	if handle, err := pcap.OpenOffline(path); err != nil {
		return err
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
	}
	return err
}

func Online() error {
	if handle, err := pcap.OpenLive("lo0", 1600, true, pcap.BlockForever); err != nil {
		return err
	} else if err := handle.SetBPFFilter("tcp and port 80"); err != nil { // optional
		return err
	} else {
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		for packet := range packetSource.Packets() {
			handlePacket(packet) // Do something with a packet here.
		}
	}
	return nil
}

func handlePacket(packet gopacket.Packet) {
	log.Print(packet)
	return
}
