package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	for _, lvl := range []struct {
		str      string
		expected Level
	}{
		{
			str:      "debug",
			expected: DebugLevel,
		},
		{
			str:      "info",
			expected: InfoLevel,
		},
		{
			str:      "warn",
			expected: WarnLevel,
		},
		{
			str:      "warning",
			expected: WarnLevel,
		},
		{
			str:      "error",
			expected: ErrorLevel,
		},
		{
			str:      "dpanic",
			expected: DPanicLevel,
		},
		{
			str:      "panic",
			expected: PanicLevel,
		},
		{
			str:      "fatal",
			expected: FatalLevel,
		},
	} {
		ll, err := ParseLevel(lvl.str)
		assert.Nil(t, err)
		assert.Equal(t, lvl.expected, ll)
	}
}
