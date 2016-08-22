package packets

import (
	"testing"
)

func TestOffline(t *testing.T) {
	err := Offline("../pcaps/mysql_mof.pcap")
	if err != nil {
		t.Error(err)
	}

}

func TestOnline(t *testing.T) {
	err := Online()
	if err != nil {
		t.Error(err)
	}
}
