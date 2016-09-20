package write

import (
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcapgo"
)

func Write(name string, snapshotLen uint32, packetSource gopacket.PacketSource) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	w := pcapgo.NewWriter(f)
	err = w.WriteFileHeader(snapshotLen, layers.LinkTypeEthernet)
	if err != nil {
		return err
	}

	c := make(chan error)
	go func() {
		for packet := range packetSource.Packets() {
			err = w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
			if err != nil {
				c <- err
			}

		}

	}()
	return <-c
}
