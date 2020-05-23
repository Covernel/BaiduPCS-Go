package converter_test

import (
	"strings"
	"testing"

	"github.com/iikira/BaiduPCS-Go/pcsutil/converter"
)

func TestTrimPathInvalidChars(t *testing.T) {
	trimed := converter.TrimPathInvalidChars("ksjadfi*/?adf")
	if strings.Compare(trimed, "ksjadfiadf") != 0 {
		t.Fatalf("trimed: %s\n", trimed)
	}
	return
}
