package packets

import (
	"log"
	"os"

	"github.com/google/gopacket/pcapgo"
)

func ReadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader, err := pcapgo.NewReader(f)
	if err != nil {
		return err
	}
	data, ci, err := reader.ReadPacketData()
	log.Print(data, ci)
	return nil
}
