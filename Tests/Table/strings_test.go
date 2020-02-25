package strings

import (
	"strings"
	"testing"
)

const msgIndex = "%s (parts: %s) - indexes: expected (%d) <> found (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text     string
		part     string
		expected int
	}{
		{"Cod3r é show", "Cod3r", 0},
		{"", "", 0},
		{"Opa", "opa", -1},
		{"leonardo", "o", 2},
		// {"leonardo", "o", 7}, // uncomment for see error
	}

	for _, test := range tests {
		t.Logf("Mass: %v", test)
		current := strings.Index(test.text, test.part)

		if current != test.expected {
			t.Errorf(msgIndex, test.text, test.part, test.expected, current)
		}
	}
}
