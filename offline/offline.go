package offline

import (
	"errors"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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
		HandlePacket(packet)
	}
	return err
}

func HandlePacket(packet gopacket.Packet) error {
	// Iterate over all layers, printing out each layer type
	for _, layer := range packet.Layers() {
		log.Print("PACKET LAYER:", layer.LayerType())
		switch layer.LayerType() {

		case layers.LayerTypeDot11InformationElement:
			handleLayerTypeDot11InformationElement(packet)

		case layers.LayerTypeIPv4:
			handleIPv4(packet)

		default:
			return errors.New("Layer not decoded.")
		}

	}
	return nil
}

func handleLayerTypeDot11InformationElement(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeDot11InformationElement).(*layers.Dot11InformationElement)
	log.Print(layer.String())
}

func handleIPv4(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeIPv4).(*layers.IPv4)
	log.Print(string(layer.Payload))
	srcDNS, _ := net.LookupAddr(layer.SrcIP.String())
	dstDNS, _ := net.LookupAddr(layer.DstIP.String())
	log.Print(srcDNS, dstDNS)

}
