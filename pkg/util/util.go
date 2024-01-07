package util

import "fmt"

func PrettyPrintByte(b byte) string {
	return fmt.Sprintf("0x%02x", b)
}
