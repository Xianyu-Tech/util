package logutil

import (
	"fmt"
	"runtime"
)

func PrintPanic() string {
	var buf [4096]byte

	cnt := runtime.Stack(buf[:], false)

	return fmt.Sprintf("%s\n", string(buf[:cnt]))
}
