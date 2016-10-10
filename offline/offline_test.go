package offline

import (
	"testing"
)

func TestOffine(t *testing.T) {
	file := "../../../../../../Desktop/2016.10.08_20-53-56-CDT.wcap"
	err := Open(file)
	if err != nil {
		t.Error(err)
	}
}
