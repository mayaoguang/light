package utils

import "fmt"

// Patch 数字补0
func Patch(i int64) string {
	return fmt.Sprintf("%08d", i)
}
