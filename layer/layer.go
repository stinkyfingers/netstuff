package layer

import (
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func Layer(packet gopacket.Packet) error {
	// Let's see if the packet is IP (even though the ether type told us)
	ipLayer := packet.Layer(layers.LayerTypeIPv4)
	if ipLayer != nil {
		log.Println("IPv4 layer detected.")
		ip, _ := ipLayer.(*layers.IPv4)

		// IP layer variables:
		// Version (Either 4 or 6)
		// IHL (IP Header Length in 32-bit words)
		// TOS, Length, Id, Flags, FragOffset, TTL, Protocol (TCP?),
		// Checksum, SrcIP, DstIP
		log.Printf("From %s to %s\n", ip.SrcIP, ip.DstIP)
		log.Println("Protocol: ", ip.Protocol)
		log.Println()
		// log.Println("PAY", string(ipLayer.LayerPayload()))
		if strings.Contains(string(ipLayer.LayerPayload()), "Trustpipe") {
			log.Print("HERE--------")
		}
	}
	return nil
}
