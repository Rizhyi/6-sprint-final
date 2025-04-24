package service

import (
	"strings"

	"github.com/rizhyi/6-sprint-final/pkg/morse"
)

func Translate(s string) string {
	// Checking what need to be translated and translate it
	if strings.ContainsAny(s, ".-") {
		result := morse.ToText(s)
		return result
	}

	result := morse.ToMorse(s)

	return result
}
