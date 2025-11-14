package string_util

type Stringer interface {
	String() string
}

// Reverse returns a string reversed.
func Reverse(s string) string {
	out := ""
	for i := len(s) - 1; i >= 0; i-- {
		out += string(s[i])
	}

	return out
}
