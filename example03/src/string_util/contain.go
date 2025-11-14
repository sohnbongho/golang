package string_util

import "strings"

func MyContain(str, substr string) bool {
	return strings.Contains(str, substr)
}
