package format

import (
	"strings"
	"testing"
)

func TestFormatChar(t *testing.T) {
	f := Format{}
	f.Initialize()
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{f.Fg.LightMagenta}, "\\[\\033[95m\\]"},
		{[]string{f.Fg.LightYellow, f.Em.Inverted}, "\\[\\033[93;7m\\]"},
	}
	for _, tt := range testCases {
		actual := FormatChar(tt.input...)
		expected := tt.expected
		if actual != expected {
			input := strings.Join(tt.input, ",")
			t.Errorf("format char failed with inputs: %s", input)
		}
	}

}

func TestEllipse(t *testing.T) {
	testCases := []struct {
		input    string
		len      int
		expected string
	}{
		{"123456789", 6, "123..."},
		{"123456", 6, "123456"},
		{"123456789", 10, "123456789"},
	}
	for _, tt := range testCases {
		actual := Ellipse(tt.input, tt.len)
		expected := tt.expected
		if actual != expected {
			t.Errorf("got: %s, want: %s", actual, expected)
		}
	}

}
