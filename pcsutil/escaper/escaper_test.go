package escaper_test

import (
	"fmt"
	"testing"

	"github.com/iikira/BaiduPCS-Go/pcsutil/escaper"
)

func TestEscape(t *testing.T) {
	fmt.Println(escaper.Escape(`asdf'asdfasd[]a[\[][sdf\[d]`, []rune{'[', '\''}))
}
