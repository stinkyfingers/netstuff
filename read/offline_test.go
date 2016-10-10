package read

import (
	"testing"
)

func TestOffline(t *testing.T) {
	file := "../../../../../../Downloads/hex2.pcap"
	res, err := Read(file)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(res))
}
