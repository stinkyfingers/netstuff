package packets

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

const (
	SNAPSHOT_LEN int32         = 1024
	TIMEOUT      time.Duration = -1 * time.Second
	PROMISCUOUS  bool          = false
)

func Write(name, device string) error {
	f, _ := os.Open(name)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		f, err = os.Create(name)
		if err != nil {
			return err
		}

	}
	// f, err := os.OpenFile(device+name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	// if err != nil {
	// 	return err
	// }
	defer f.Close()

	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(SNAPSHOT_LEN), layers.LinkTypeEthernet)

	// Open device
	handle, err = pcap.OpenLive(device, SNAPSHOT_LEN, PROMISCUOUS, TIMEOUT)
	if err != nil {
		return err
	}
	defer handle.Close()

	packetCount := 0

	// Start processing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		fmt.Println("++++", packet)
		err = w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
		if err != nil {
			return err
		}

		packetCount++
		log.Print(packetCount)

		// Only capture 100 and then stop
		if packetCount > 100 {
			break
		}
	}
	return nil

}
