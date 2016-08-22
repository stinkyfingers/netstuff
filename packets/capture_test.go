package packets

import (
	"log"
	"net"
	"os"
	"testing"
)

func TestCapture(t *testing.T) {
	ifaces, err := net.Interfaces()
	if err != nil {
		t.Error(err)
	}
	for _, iface := range ifaces {
		log.Printf("Trying capture on %q", iface.Name)
		if err := Capture(iface); err != nil {
			log.Printf("Error capturing on %q: %v", iface.Name, err)
		} else {
			log.Printf("Successfully captured on %q", iface.Name)
			return
		}
	}
	os.Exit(1)
}
