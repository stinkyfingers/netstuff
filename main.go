package main

//http://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
import (
	"flag"
	"log"

	"github.com/stinkyfingers/netstuff/device"
	"github.com/stinkyfingers/netstuff/offline"
	"github.com/stinkyfingers/netstuff/packet"
)

var (
	file        = flag.String("pcap", "", "--pcap=filename; needed for offline processing of pcaps")
	listDevices = flag.Bool("devices", false, "--devices=true; show devices")
	listen      = flag.String("listen", "", "--listen=<device_name>")
)

func main() {
	var err error
	flag.Parse()

	//device
	if *listDevices {
		devices, err := device.Devices()
		if err != nil {
			log.Fatal(err)
		}
		device.Print(devices)
	}

	if *listen != "" {
		packetSource, err := device.OpenDevice(*listen)
		if err != nil {
			log.Fatal(err)
		}
		err = packet.HandlePackets(packetSource)
		if err != nil {
			log.Fatal(err)
		}
	}

	//offline
	if *file != "" {
		err = handleOffline(*file)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	log.Print("DONE")
}

func handleOffline(filename string) error {
	return offline.Open(filename)
}
