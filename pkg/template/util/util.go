package util

import (
	"fmt"
	"strings"
	"unicode"
)

func FirstLower(input string) string {
	for i, v := range input {
		return string(unicode.ToLower(v)) + input[i+1:]
	}
	return ""
}

func FirstUpper(input string) string {
	for i, v := range input {
		return string(unicode.ToUpper(v)) + input[i+1:]
	}
	return ""
}

func Batch(input string) string {
	if strings.HasSuffix(input, "s") {
		return fmt.Sprintf("%ses", input)
	}
	return fmt.Sprintf("%ss", input)
}

func IF[T any](b bool, this T, that T) T {
	if b {
		return this
	} else {
		return that
	}
}
