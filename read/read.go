package read

import (
	"bytes"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Read(name string) ([]byte, error) {
	var buffer bytes.Buffer
	handle, err := pcap.OpenOffline(name)
	if err != nil {
		return buffer.Bytes(), err
	}
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		buffer.Write(packet.Data())
	}
	return buffer.Bytes(), err
}
