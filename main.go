package main

//http://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
import (
	"flag"
	"log"

	"github.com/stinkyfingers/netstuff/device"
	"github.com/stinkyfingers/netstuff/offline"
)

var (
	file        = flag.String("pcap", "", "--pcap=filename; needed for offline processing of pcaps")
	listDevices = flag.Bool("devices", false, "--devices=true; show devices")
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
