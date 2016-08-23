package main

//http://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
import (
	"errors"
	"flag"
	"log"
	"net"

	"github.com/stinkyfingers/netstuff/packets"
)

var (
	offline  = flag.Bool("off", false, "--off=true if you want to do offline pcap processing")
	filename = flag.String("f", "", "-f=filename; needed for offline processing of pcaps")
)

func main() {
	var err error
	flag.Parse()

	//offline
	if *offline {
		err = handleOffline(*filename)
		if err != nil {
			log.Print(err)
		}
		return
	}

	//online
	err = handleOnline(*filename)
	if err != nil {
		log.Print(err)
	}

	log.Print("DONE")
}

func handleOffline(filename string) error {
	if filename == "" {
		return errors.New("No filename; -f flag")
	}
	return packets.Offline(filename)
}

func handleOnline(filename string) error {
	name := "test.pcap"
	if filename != "" {
		name = filename
	}
	// _ = packets.Devices()

	ifaces, err := net.Interfaces()
	if err != nil {
		return err
	}
	for _, iface := range ifaces {
		// addresses, err := iface.Addrs()
		// name := iface.Name
		// log.Print("ADDRESS & NAME", addresses, err, name)
		// err = packets.OpenLive(iface)
		// if err != nil {
		// 	log.Print(err)
		// }

		err = packets.Write(name, iface.Name)
		if err != nil {
			log.Print("ERR", err)
		}
	}
	return err
}
