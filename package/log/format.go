package log

import (
	"fmt"
	"strings"
)

type Format int

const (
	FormatConsole Format = iota
	FormatJSON
)

func ParseFormat(format string) (Format, error) {
	switch strings.ToLower(format) {
	case "console":
		return FormatConsole, nil
	case "json":
		return FormatJSON, nil
	}

	return FormatConsole, fmt.Errorf("not a valid Format: %q", format)
}
