package piscine

import (
	"strings"
)

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	var builder strings.Builder
	for i := 2; i < len(str); i += 3 {
		builder.WriteByte(str[i])
	}
	if builder.Len() == 0 {
		return "\n"
	}

	return builder.String() + "\n"
}
