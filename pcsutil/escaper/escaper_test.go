package escaper_test

import (
	"fmt"
	"github.com/Denis-Oakley/BaiduPCS-Go/pcsutil/escaper"
	"testing"
)

func TestEscape(t *testing.T) {
	fmt.Println(escaper.Escape(`asdf'asdfasd[]a[\[][sdf\[d]`, []rune{'[', '\''}))
}
