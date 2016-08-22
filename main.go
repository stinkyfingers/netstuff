package main

//http://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
import (
	"log"
	"net"

	"github.com/stinkyfingers/netstuff/packets"
)

func main() {
	// packets.Offline("./pcaps/mysql_mof.pcap")
	ifaces, err := net.Interfaces()
	log.Print(ifaces)
	if err != nil {
		log.Fatal(err)
	}
	for _, iface := range ifaces {
		err = packets.Capture(iface)
		if err != nil {
			log.Fatal(err)
		}
	}

	log.Print("DONE")
}
