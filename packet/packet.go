package packet

import (
	"errors"
	"log"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func HandlePackets(packetSource *gopacket.PacketSource) error {
	for packet := range packetSource.Packets() {
		err := HandlePacket(packet)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandlePacket(packet gopacket.Packet) error {
	layer := packet.Layer(layers.LayerTypeIPv4)
	if layer == nil {
		return errors.New("Content not IPv4")
	}
	ipv4Layer, ok := layer.(*layers.IPv4)
	if !ok {
		return errors.New("Failed to case IP layer to IPv4")
	}
	transportLayerType := ipv4Layer.NextLayerType()
	transportLayer := packet.Layer(transportLayerType)
	var dstPort, srcPort int

	switch transportLayerType {
	case layers.LayerTypeTCP:
		if transportLayer == nil {
			return errors.New("Packet reported as TCP but TCP layer is nil")
		}
		srcPort = int(transportLayer.(*layers.TCP).SrcPort)
		dstPort = int(transportLayer.(*layers.TCP).DstPort)

	case layers.LayerTypeUDP:
		if transportLayer == nil {
			return errors.New("Packet reported as UDP but UDP layer is nil")
		}
		srcPort = int(transportLayer.(*layers.UDP).SrcPort)
		dstPort = int(transportLayer.(*layers.UDP).DstPort)

	case layers.LayerTypeSCTP:
		if transportLayer == nil {
			return errors.New("Packet reported as SCTP but SCTP layer is nil")
		}
		srcPort = int(transportLayer.(*layers.SCTP).SrcPort)
		dstPort = int(transportLayer.(*layers.SCTP).DstPort)
	}

	log.Print(ipv4Layer.Payload)
	log.Print(srcPort, dstPort)
	return nil
}

// TODO - do I have use for iterating over all layers?
func HandlePacketLayers(packet gopacket.Packet) error {
	// Iterate over all layers, printing out each layer type
	for _, layer := range packet.Layers() {
		log.Print("PACKET LAYER: ", layer.LayerType())
		switch layer.LayerType() {

		case layers.LayerTypeDot11InformationElement:
			handleLayerTypeDot11InformationElement(packet)

		case layers.LayerTypeIPv4:
			handleIPv4(packet)

		case layers.LayerTypeTCP:
			handleTCP(packet)

		default:
			// return errors.New("Layer not decoded.")
			log.Print("Layer not detected")
			continue
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
	log.Print("IPS: ", srcDNS, dstDNS)
}

func handleTCP(packet gopacket.Packet) {
	layer := packet.Layer(layers.LayerTypeTCP).(*layers.TCP)
	log.Print(string(layer.Payload))
	srcDNS, _ := net.LookupAddr(layer.SrcPort.String())
	dstDNS, _ := net.LookupAddr(layer.DstPort.String())
	log.Print("PORTS: ", srcDNS, dstDNS)
}
