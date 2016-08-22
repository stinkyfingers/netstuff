package packets

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/google/gopacket/pcap"
)

func Capture(iface net.Interface) error {
	var h *pcap.Handle
	var err error

	h, err = pcap.OpenLive(iface.Name, 65536, false, time.Second*3)
	if err != nil {
		return err
	}
	go generatePackets()
	h.ReadPacketData()
	data, ci, err := h.ReadPacketData()
	if err != nil {
		return fmt.Errorf("readpacketdata: %v", err)
	}
	log.Printf("Read packet, %v bytes, CI: %+v", len(data), ci)
	return nil

}

func generatePackets() {
	if resp, err := http.Get("http://code.google.com"); err != nil {
		log.Printf("Could not get HTTP: %v", err)
	} else {
		log.Print("called google.com")
		resp.Body.Close()
	}
}
